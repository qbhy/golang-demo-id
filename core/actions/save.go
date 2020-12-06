package actions

type SaveAction struct {
	Action
}

func (action SaveAction) Handle(data DataMap) {
	data[action.Key] = action.Arg

	action.ReadChan <- true
}

func (action SaveAction) IsWrite() bool {
	return true
}
