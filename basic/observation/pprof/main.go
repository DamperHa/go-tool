package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
)

// 注册了以下路由
/*
func init() {
	http.HandleFunc("/debug/pprof/", Index)
	http.HandleFunc("/debug/pprof/cmdline", Cmdline)
	http.HandleFunc("/debug/pprof/profile", Profile)
	http.HandleFunc("/debug/pprof/symbol", Symbol)
	http.HandleFunc("/debug/pprof/trace", Trace)
}
*/

func main() {
	// 注册路由
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(*request)
		writer.Write([]byte("hello"))
	})

	// 服务对象
	s := http.Server{
		Addr: "localhost:8080",
	}

	// 优雅退出
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-c
		s.Shutdown(context.Background())
	}()

	log.Println(s.ListenAndServe())
}
