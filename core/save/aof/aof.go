package aof

import (
	"bufio"
	"fmt"
	"id/core/contracts"
	"id/util"
	"io"
	"os"
	"strings"
	"time"
)

type Aof struct {
	Path          string                    // 文件目录
	Type          string                    // 类型, 支持实时和每秒
	ActionsBuffer []contracts.CommandAction // 缓存一秒的 action
	file          *os.File                  // 文件句柄
}

const (
	REALTIME = "realtime"
	SECONDS  = "seconds"
)

func (aof *Aof) Init() {
	aof.file = util.OpenFileOrCreate(aof.Path)
}

func (aof *Aof) Run(actionChan chan contracts.CommandAction) {
	/**
	0. 根据持久化规则判断是否满足持久化条件，若满足则往下执行
	1. 把新的 action 写入缓存取，每隔一定时间把缓存取的日志写入文件/或者实时写入文件(影响性能)
	2. 检查日志文件是否达到需要重新的规模，若满足，则执行重写
	*/
	// 开启协程监听持久化通道
	go func() {
		if aof.Type == REALTIME { // 实时写入
			for action := range actionChan {
				aof.Save(action.Data, action)
			}
		} else { // 定时写入

			go func() { // 秒级定时器
				ticker := time.NewTicker(time.Second)
				for {
					<-ticker.C
					if len(aof.ActionsBuffer) > 0 { // 缓存中有待写入的数据，就执行一次写入
						actions := aof.ActionsBuffer
						aof.ActionsBuffer = []contracts.CommandAction{}
						aof.BatchSave(actions)
					}
				}
			}()

			for action := range actionChan {
				aof.ActionsBuffer = append(aof.ActionsBuffer, action)
			}
		}
	}()
}

func (aof *Aof) BatchSave(actions []contracts.CommandAction) {
	content := ""
	for _, action := range actions {
		content += fmt.Sprintf("%s,%s,%d\n", action.Command.Id(), action.Key, action.Arg)
	}

	_, err := aof.file.Write([]byte(content))
	if err != nil {
		fmt.Println("写文件失败", err)
	}
}

func (aof *Aof) Save(data contracts.DataMap, action contracts.CommandAction) {
	_, err := aof.file.Write([]byte(fmt.Sprintf("%s,%s,%d\n", action.Command.Id(), action.Key, action.Arg)))
	if err != nil {
		fmt.Println(err)
	}
}

func (aof *Aof) Recovery(center contracts.DataCenter) {
	br := bufio.NewReader(aof.file)
	center.SetRecovering(true)
	for {
		line, _, c := br.ReadLine()

		if c == io.EOF {
			break
		}


		actionArr := strings.Split(string(line), ",")
		if len(actionArr) == 3 { // 长度不对说明数据有问题
			center.Call(actionArr[0], actionArr[1], util.StrToInt(actionArr[2]))
		}
	}
	center.SetRecovering(false)
}
