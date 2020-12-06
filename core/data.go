package core

import (
	"id/core/actions"
	"sync"
)

// 定义数据存储中心
type DataCenter struct {
	actionChan  chan actions.ActionAbility
	data        actions.DataMap
	savable     string                     // 持久化方案
	savableChan chan actions.ActionAbility // 通知持久化的通道
	savableLock *sync.RWMutex
}

// 设置值
func (dc DataCenter) Set(key string, value int) bool {
	action := actions.NewAction(key, value)

	dc.actionChan <- actions.SaveAction{Action: action} // 使用 save action
	return (<-action.ReadChan).(bool)
}

// 获取值
func (dc DataCenter) Get(key string) int {
	action := actions.NewAction(key, 0)

	dc.actionChan <- actions.GetAction{Action: action} // 使用 get action
	return (<-action.ReadChan).(int)
}

// 自增
func (dc DataCenter) Incr(key string, num int) int {
	action := actions.NewAction(key, num)

	dc.actionChan <- actions.IncrAction{Action: action} // 使用自增 action
	return (<-action.ReadChan).(int)
}

// ... 更多 action 绑定

// new 一个数据中心
func NewData() DataCenter {
	return DataCenter{
		actionChan:  make(chan actions.ActionAbility),
		data:        actions.DataMap{},
		savable:     "", // 默认不开启持久化
		savableChan: nil,
		savableLock: new(sync.RWMutex),
	}
}

// 开始监听
func (data DataCenter) Start() {
	go func() {
		for action := range data.actionChan {
			if action.IsWrite() && data.savableChan != nil { // 写入操作的话，可能需要持久化
				data.savableLock.Lock()
				action.Handle(data.data)
				data.savableLock.Unlock()

				data.savableChan <- action
			} else {
				action.Handle(data.data)
			}
			// 思考：非持久化情况下是否需要加个读写锁 ？
		}
	}()
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
func (data *DataCenter) Savable(method string) {

	data.savable = method
	data.savableChan = make(chan actions.ActionAbility)

	// 开启协程监听持久化通道
	go func() {
		for range data.savableChan {
			if method == "规则1" {
				/**
				0. 根据持久化规则判断是否满足持久化条件，若满足则往下执行
				1. 使用读写锁锁住写操作
				2. 复制一个 data map
				3. 解开读写锁
				4. 把复制好的 data map 里面的数据序列化成字符串写入指定文件中
				*/
			} else if method == "规则2" {
				/**
				0. 根据持久化规则判断是否满足持久化条件，若满足则往下执行
				1. 把新的 action 写入缓存取，没隔一定时间把缓存取的日志写入文件/或者实时写入文件(影响性能)
				2. 检查日志文件是否达到需要重新的规模，若满足，则执行重写
				*/
			} else if method == "规则3" {
				/**
				规则3 略
				*/
			}
		}
	}()

}
