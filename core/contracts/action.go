package contracts

type CommandAction struct {
	Command  ExecuteAble
	Key      string
	Arg      int
	BackChan chan interface{}
	Data     DataMap
}

type ExecuteAble interface {
	IsWrite() bool
	Id() string
	Execute(dataMap DataMap, key string, args int, ch chan interface{})
}
