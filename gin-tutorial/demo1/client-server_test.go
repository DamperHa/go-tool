package demo1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html"
	"log"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 创建一个Foo路由和处理函数
	// http.Handle("/foo", fooHandler)

	// 创建一个bar路由和处理函数
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	// 监听8080端口
	log.Fatal(http.ListenAndServe(":8080", nil))

	r.Run() // listen and serve on 0.0.0.0:8080
}
