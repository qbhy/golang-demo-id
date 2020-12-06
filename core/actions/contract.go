package actions

type DataMap map[string]int // 改成任意类型加上自定义的一些 action ，理论上可以实现类似 redis 的 set、map、list等更丰富的功能

type ActionAbility interface {
	Handle(dataMap DataMap)
	IsWrite() bool
}
