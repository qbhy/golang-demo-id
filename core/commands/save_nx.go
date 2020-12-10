package commands

import (
	"id/core/contracts"
)

type SaveNxCommand struct {
}

const SAVE_NX_COMMAND = "savenx"

func (command SaveNxCommand) Id() string {
	return SAVE_NX_COMMAND
}

func (command SaveNxCommand) Execute(data contracts.DataMap, key string, args int, ch chan interface{}) {
	_, ok := data[key]
	if ok { // 存在
		ch <- false
		return
	}

	data[key] = args

	ch <- true
}

func (command SaveNxCommand) IsWrite() bool {
	return true
}
