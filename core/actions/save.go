package actions

import "id/core/contracts"

type SaveAction struct {
	Action
}

func (action SaveAction) Handle(data contracts.DataMap) {
	data[action.Key] = action.Arg

	action.ReadChan <- true
}

func (action SaveAction) IsWrite() bool {
	return true
}
