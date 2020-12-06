package actions

import "id/core/contracts"

type ExistsAction struct {
	Action
}

func (action ExistsAction) Handle(data contracts.DataMap) {
	_, ok := data["key1"]

	action.ReadChan <- ok
}

func (action ExistsAction) IsWrite() bool {
	return false
}
