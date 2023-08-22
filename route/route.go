package route

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/api/someendpoint", handlers.SomeHandler)
	// 添加其他路由和处理函数

	return r
}
