package actions

type GetAction struct {
	Action
}

func (action GetAction) Handle(data DataMap) {
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
