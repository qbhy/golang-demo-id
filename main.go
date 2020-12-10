package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"id/core"
	"id/core/contracts"
	"id/core/save/aof"
	"id/core/save/rdb"
	"id/core/save/remix"
	"id/util"
	"net/http"
	"os"
)

var data contracts.DataCenter

/**
要实现高可用性，还得实现以下功能：
1. 主从复制
2. 支持集群架构，可以参考 redis 槽点设计
*/
func main() {
	_ = godotenv.Load(".env")

	data = core.NewData()

	aofInstance := &aof.Aof{
		Path: os.Getenv("AOF_FILE_PATH"),
		Type: os.Getenv("AOF_TYPE"),
	}
	rdbInstance := &rdb.Rdb{
		Path: os.Getenv("RDB_FILE_PATH"),
	}
	remixInstance := &remix.Remix{
		Path: os.Getenv("REMIX_FILE_PATH"),
	}

	// 开启持久化协程，开启使用 aof 持久化方案
	savable := map[string]contracts.Savable{
		contracts.AOF:   aofInstance,
		contracts.RDB:   rdbInstance,
		contracts.REMIX: remixInstance,
	}

	switch savableType := os.Getenv("SAVE_TYPE"); savableType {
	case contracts.AOF, contracts.RDB, contracts.REMIX:
		data.Savable(savable[savableType])
	}

	// 开启操作管道协程
	data.Start()

	router := gin.Default()


	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world")
	})

	// 对外调用做一些映射，后面版本会改成 tcp
	router.POST("/cmd", func(context *gin.Context) {
		ch := data.Call(context.PostForm("cmd"), context.PostForm("key"), util.StrToInt(context.PostForm("arg")))
		var result interface{} = nil
		if ch != nil {
			result = <-ch
		}
		context.JSON(http.StatusOK, map[string]interface{}{
			"result": result,
		})
	})

	_ = router.Run(":8001")

}
