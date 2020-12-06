package actions

import "id/core/contracts"

type DeleteAction struct {
	Action
}

func (action DeleteAction) Handle(data contracts.DataMap) {
	delete(data, action.Key)
	action.ReadChan <- true
}
func (action DeleteAction) IsWrite() bool {
	return true
}
