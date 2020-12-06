package actions

type ExistsAction struct {
	Action
}

func (action ExistsAction) Handle(data DataMap) {
	_, ok := data["key1"]

	action.ReadChan <- ok
}

func (action ExistsAction) IsWrite() bool {
	return false
}
