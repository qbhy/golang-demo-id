package commands

import (
	"id/core/contracts"
)

type ExistsCommand struct {
}

const EXISTS_COMMAND = "exists"

func (action ExistsCommand) Id() string {
	return EXISTS_COMMAND
}

func (command ExistsCommand) Execute(data contracts.DataMap, key string, args int, ch chan interface{}) {

	_, ok := data[key]

	ch <- ok
}

func (action ExistsCommand) IsWrite() bool {
	return false
}
