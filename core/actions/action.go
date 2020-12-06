package actions

// 定一个一个操作
type Action struct {
	Key      string           // 要操作key
	Arg      int              // 参数
	ReadChan chan interface{} // 返回数据的通道
}

func NewAction(key string, value int) Action {
	return Action{
		Key:      key,
		Arg:      value,
		ReadChan: make(chan interface{}),
	}
}
