package actions

type DeleteAction struct {
	Action
}

func (action DeleteAction) Handle(data DataMap) {
	delete(data, action.Key)
	action.ReadChan <- true
}
func (action DeleteAction) IsWrite() bool {
	return true
}
