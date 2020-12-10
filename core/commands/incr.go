package commands

import (
	"id/core/contracts"
)

type IncrCommand struct {
}

const INCR_COMMAND = "incr"

func (command IncrCommand) Id() string {
	return INCR_COMMAND
}

func (command IncrCommand) Execute(data contracts.DataMap, key string, value int, ch chan interface{}) {
	num, ok := data[key]

	if ok {
		data[key] = num + value
	} else {
		data[key] = value
	}

	ch <- data[key]
}

func (command IncrCommand) IsWrite() bool {
	return true
}
