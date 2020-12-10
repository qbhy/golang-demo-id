package contracts

import "sync"

type DataMap map[string]int // 改成任意类型加上自定义的一些 action ，理论上可以实现类似 redis 的 set、map、list等更丰富的功能

type DataCenter interface {
	Commands(commands map[string]ExecuteAble) DataCenter       // 绑定多个命令
	Bind(key string, command ExecuteAble) DataCenter           // 绑定单个命令
	Call(command string, key string, arg int) chan interface{} // 调用某命令
	Start()                                                    // 开始监听 command
	Savable(savable Savable)                                   // 开启持久化
	WriteLock() *sync.RWMutex                                  // 获取写锁
	GetDataMap() *DataMap
	SetRecovering(Recovering bool)
}
