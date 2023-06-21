// Package net @Author:冯铁城 [17615007230@163.com] 2023-06-19 17:22:41
package net

import (
	"log"
	"net/http"
	"strings"
)

// StartHttpServer 开启服务器
func StartHttpServer() {

	//1.加载Handler
	http.HandleFunc("/headers", getHeadHandler)
	http.HandleFunc("/path-params/", getPathParamHandler)
	http.HandleFunc("/url-params", getUrlParamHandler)

	//2.开启服务器并监听,如果出现异常,进行异常打印
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

// 获取请求头handler
func getHeadHandler(w http.ResponseWriter, req *http.Request) {

	//1.获取单个请求头
	token := req.Header.Get("token")
	log.Printf("request header token is [%v]\n", token)

	//2.获取所有请求头
	headers := req.Header
	for key := range headers {
		log.Printf("request header %v is %v\n", key, headers[key])
	}
}

// 获取路径参数handler
func getPathParamHandler(w http.ResponseWriter, req *http.Request) {

	//1.获取请求路由
	route := req.URL.Path[1:]
	log.Printf("request route is [%v]\n", route)

	//2.获取路由参数
	pathParam := route[strings.LastIndex(route, "/")+1:]
	log.Printf("request path param is [%v]\n", pathParam)
}

// 获取url参数handler
func getUrlParamHandler(w http.ResponseWriter, req *http.Request) {

	//1.获取单个URL参数
	age := req.URL.Query().Get("age")
	log.Printf("request url param age is [%v]\n", age)

	//2.获取所有URL参数
	query := req.URL.Query()
	for key := range query {
		log.Printf("request url param %v is %v\n", key, query[key])
	}
}

func getBodyJsonParamHandler(w http.ResponseWriter, req *http.Request) {

}

//// HelloHandler
//func helloHandler(w http.ResponseWriter, req *http.Request) {
//
//	//4.响应
//	fmt.Fprintf(w, "Hello! "+route[strings.LastIndex(route, "/")+1:])
//}
