package actions

import "id/core/contracts"

type GetAction struct {
	Action
}

func (action GetAction) Handle(data contracts.DataMap) {
	num, ok := data[action.Key]

	if ok {
		action.ReadChan <- num
	} else {
		action.ReadChan <- action.Arg
	}

}


func (action GetAction) IsWrite() bool {
	return false
}
