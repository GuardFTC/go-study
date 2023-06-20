// Package net @Author:冯铁城 [17615007230@163.com] 2023-06-19 17:22:41
package net

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// StartHttpServer 开启服务器
func StartHttpServer() {

	//1.指定路由
	http.HandleFunc("/api/nets/", httpServer)

	//2.开启服务器并监听
	err := http.ListenAndServe("127.0.0.1:8080", nil)

	//3.如果出现异常,进行异常打印
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

// HttpServer 定义http服务器
func httpServer(w http.ResponseWriter, req *http.Request) {

	//1.获取路由参数
	route := req.URL.Path[1:]

	//2.响应
	fmt.Fprintf(w, "Hello! "+route[strings.LastIndex(route, "/")+1:])
}
