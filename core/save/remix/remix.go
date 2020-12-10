package remix

import (
	"bufio"
	"fmt"
	"id/core/contracts"
	"id/util"
	"io"
	"os"
)

type Remix struct {
	Path string // 文件目录
	file *os.File
}

func (remix Remix) Init() {
	remix.file = util.OpenFileOrCreate(remix.Path)
}

func (remix Remix) Run(actionChan chan contracts.CommandAction) {
	/**
	0. 根据持久化规则判断是否满足持久化条件，若满足则往下执行
	1. 使用读写锁锁住写操作
	2. 复制一个 data map
	3. 解开读写锁
	4. 把复制好的 data map 里面的数据序列化成字符串写入指定文件中
	*/
	// todo: 待实现
}

func (remix Remix) Save(data contracts.DataMap, action contracts.CommandAction) {
	content := ""
	for key, value := range data {
		content += fmt.Sprintf("%s,%d\n", key, value)
	}

	_, err := remix.file.Write([]byte(content))

	if err != nil {
		fmt.Println(err)
	}
}

func (remix *Remix) Recovery(center contracts.DataCenter) {
	br := bufio.NewReader(remix.file)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		fmt.Println(string(a))
	}
}