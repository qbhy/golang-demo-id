package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"id/core"
	"net/http"
)

var data core.DataCenter

func main() {
	data = core.NewData()

	// 开启持久化协程
	data.Savable("规则2")

	// 开启操作管道协程
	data.Start()

	data.Set("a", 1)

	fmt.Println(data.Incr("a", 100))

	router := gin.Default()

	// 对外调用做一些映射

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world")
	})

	_ = router.Run(":8001")

}
