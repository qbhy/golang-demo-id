package commands

import (
	"id/core/contracts"
)

type SaveCommand struct {
}

const SAVE_COMMAND = "save"

func (command SaveCommand) Id() string {
	return SAVE_COMMAND
}

func (command SaveCommand) Execute(data contracts.DataMap, key string, args int, ch chan interface{}) {
	data[key] = args

	ch <- true
}

func (command SaveCommand) IsWrite() bool {
	return true
}
