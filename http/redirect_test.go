package http

import (
	"net/http"
	"testing"
)

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	// 设置HTTP Status Code为301 Moved Permanently，并重定向到/new-page路径
	http.Redirect(w, r, "/new-page", http.StatusMovedPermanently)
}

func newPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the new page!"))
}

func TestRedirect(t *testing.T) {
	// 当访问/old-page时，调用redirectHandler进行重定向
	http.HandleFunc("/old-page", redirectHandler)
	// 当访问/new-page时，调用newPageHandler展示新页面内容
	http.HandleFunc("/new-page", newPageHandler)

	// 监听8080端口
	http.ListenAndServe(":8080", nil)
}
