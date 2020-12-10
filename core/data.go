package core

import (
	"id/core/commands"
	"id/core/contracts"
	"sync"
)

// 定义数据存储中心
type Data struct {
	commandChan chan contracts.CommandAction
	data        contracts.DataMap
	savable     contracts.Savable            // 持久化方案
	savableChan chan contracts.CommandAction // 通知持久化的通道
	Lock        *sync.RWMutex                // 写锁
	Recovering  bool                         // 是否处于恢复中的状态
	commands    map[string]contracts.ExecuteAble
}

// 批量注册 Command
func (data *Data) Commands(commands map[string]contracts.ExecuteAble) contracts.DataCenter {
	data.commands = commands

	return data
}

// 绑定单个 Command
func (data *Data) Bind(key string, command contracts.ExecuteAble) contracts.DataCenter {
	data.commands[key] = command
	return data
}

// 调用命令
func (data *Data) Call(command string, key string, arg int) chan interface{} {
	ability, ok := data.commands[command]

	if ok {
		ch := make(chan interface{}, 1)

		if data.Recovering { // 恢复模式下直接执行即可
			ability.Execute(data.data, key, arg, ch)
			return nil
		}

		data.commandChan <- contracts.CommandAction{
			Key:      key,
			Arg:      arg,
			Command:  ability,
			BackChan: ch,
			Data:     data.data,
		}
		return ch
	}

	return nil
}

// new 一个数据中心
func NewData() contracts.DataCenter {
	return (&Data{
		commandChan: make(chan contracts.CommandAction),
		data:        contracts.DataMap{},
		savable:     nil, // 默认不开启持久化
		savableChan: nil,
		Lock:        new(sync.RWMutex),
	}).Commands(map[string]contracts.ExecuteAble{
		commands.DELETE_COMMAND:  commands.DeleteCommand{},
		commands.GET_COMMAND:     commands.GetCommand{},
		commands.EXISTS_COMMAND:  commands.ExistsCommand{},
		commands.INCR_COMMAND:    commands.IncrCommand{},
		commands.SAVE_COMMAND:    commands.SaveCommand{},
		commands.SAVE_NX_COMMAND: commands.SaveNxCommand{},
	})
}

// 开始监听
func (data Data) WriteLock() *sync.RWMutex {
	return data.Lock
}

func (data Data) Start() {
	go func() {
		for action := range data.commandChan {
			if action.Command.IsWrite() && data.savableChan != nil { // 写入操作的话，可能需要持久化
				data.WriteLock().Lock()
				action.Command.Execute(data.data, action.Key, action.Arg, action.BackChan)
				data.WriteLock().Unlock()

				data.savableChan <- action
			} else {
				action.Command.Execute(data.data, action.Key, action.Arg, action.BackChan)
			}
			// todo: 思考：非持久化情况下是否需要加个读写锁 ？
		}
	}()
}

func (data *Data) GetDataMap() *contracts.DataMap {
	return &data.data
}

func (data *Data) SetRecovering(Recovering bool) {
	data.Recovering = Recovering
}

/**
开启持久化
	原理：持久化原理就是将内存中的数据写到文件中
	持久化思路可以参考 Redis 持久化的两个思路：
		1. 按照规则执行全量持久化（比如1分钟内写操作达到1000次执行一次持久化），需要读写锁
		2. 记录每次写操作的日志，重启的时候根据日志重新执行写操作（记录太多，可能需要定时重写日志）
		3. 混合持久化，文件前面一部分放数据，后面放写操作日志

	以下是持久化数据的伪代码
*/
func (data *Data) Savable(savable contracts.Savable) {
	data.savable = savable
	data.savableChan = make(chan contracts.CommandAction)
	data.savable.Init()

	// 从持久化文件中恢复数据
	data.savable.Recovery(data)

	data.savable.Run(data.savableChan)

}
