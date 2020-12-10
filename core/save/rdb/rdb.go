package rdb

import (
	"fmt"
	"id/core/contracts"
	"id/util"
	"os"
)

type Rdb struct {
	Path string // 文件目录
	file *os.File
}

func New() Rdb {
	return Rdb{}
}

func (rdb Rdb) Init() {
	rdb.file = util.OpenFileOrCreate(rdb.Path)
}

func (rdb Rdb) Run(actionChan chan contracts.CommandAction) {
	/**
	0. 根据持久化规则判断是否满足持久化条件，若满足则往下执行
	1. 使用读写锁锁住写操作
	2. 复制一个 data map
	3. 解开读写锁
	4. 把复制好的 data map 里面的数据序列化成字符串写入指定文件中
	*/
}

func (rdb Rdb) Save(data contracts.DataMap, action contracts.CommandAction) {
	content := ""
	for key, value := range data {
		content += fmt.Sprintf("%s,%d\n", key, value)
	}

	_, err := rdb.file.Write([]byte(content))

	if err != nil {
		fmt.Println(err)
	}
}
