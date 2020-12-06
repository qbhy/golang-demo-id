package actions

import "id/core/contracts"

type IncrAction struct {
	Action
}

func (action IncrAction) Handle(data contracts.DataMap) {
	num, ok := data[action.Key]

	if ok {
		data[action.Key] = num + action.Arg
	} else {
		data[action.Key] = action.Arg
	}

	action.ReadChan <- data[action.Key]
}


func (action IncrAction) IsWrite() bool {
	return true
}
