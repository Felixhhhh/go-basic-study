package main

import (
	"fmt"
	"net/http"
)

// 1、定义处理器函数
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	// 请求方式：GET POST DELETE PUT UPDATE
	fmt.Println("method:", r.Method)
	// /go
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)
	// 回复
	w.Write([]byte("www.5lmh.com"))

}

func main() {
	//http://127.0.0.1:8000/go
	// 单独写回调函数
	// 2、绑定 路由"/hello"和处理器函数helloHandler
	http.HandleFunc("/go", Handler)
	//http.HandleFunc("/ungo",myHandler2 )
	// addr：监听的地址
	// handler：回调函数
	// 3、启动服务并监听8080端口
	http.ListenAndServe("127.0.0.1:8000", nil)
}
