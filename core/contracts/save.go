package contracts

const (
	RDB   = "rdb"
	AOF   = "aof"
	REMIX = "remix"
)

type Savable interface {
	Init()                                   // 初始化
	Run(actionChan chan CommandAction)       // 开启持久化监听
	Save(data DataMap, action CommandAction) // 执行持久化
	Recovery(center DataCenter)
}
