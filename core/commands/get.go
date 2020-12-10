package commands

import (
	"id/core/contracts"
)

type GetCommand struct {
}

const GET_COMMAND = "get"

func (action GetCommand) Id() string {
	return GET_COMMAND
}

func (command GetCommand) Execute(data contracts.DataMap, key string, args int, ch chan interface{}) {
	num, ok := data[key]

	if ok {
		ch <- num
	} else {
		ch <- args
	}
}

func (command GetCommand) IsWrite() bool {
	return false
}
