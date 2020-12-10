package commands

import (
	"id/core/contracts"
)

type DeleteCommand struct {
}

const DELETE_COMMAND = "delete"

func (command DeleteCommand) Id() string {
	return DELETE_COMMAND
}

func (command DeleteCommand) Execute(data contracts.DataMap, key string, args int, ch chan interface{}) {
	delete(data, key)
	ch <- true
}

func (command DeleteCommand) IsWrite() bool {
	return true
}
