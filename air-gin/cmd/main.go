package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个 Gin 路由实例
	r := gin.Default()

	r.Use(func(c *gin.Context) {})
	// 放在这个目录下他能识别吗
	// 定义一个处理 GET 请求的路由
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Gin ok nihao ffaff faf  lainghao zhenshua  zhenhaoyongi!")
	})

	// 启动服务器，监听在 8080 端口
	r.Run(":8089")
}
