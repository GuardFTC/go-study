// Package net @Author:冯铁城 [17615007230@163.com] 2023-06-20 10:06:57
package net

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/thedevsaddam/gojsonq/v2"
)

const formData = "<p class=\"\\\">1.点击用户头像设置</p>\n<p class=\"img\">\n\t<img style=\"width: 50%\" src=\"https://ecosys-service.cicic.cn/blob/ecosys/public/resource/content/pic/20230602/168569497215980262.jpeg\"/>\n</p>\n<p class=\"\\\">2.设置用户头像</p>\n<p class=\"img\">\n\t<img style=\"width: 50%\" src=\"https://ecosys-service.cicic.cn/blob/ecosys/public/resource/content/pic/20230602/168569498519565620.jpeg\"/>\n</p>\n<p class=\"\\\">设置完毕</p>"

// StartComplexHttpServer 开启Http复杂服务器
func StartComplexHttpServer() {

	//1.加载处理器
	http.HandleFunc("/hello", initHandler(helloWorldHandler))
	http.HandleFunc("/complex", initHandler(complexHandler))

	//2.开启服务器
	err := http.ListenAndServe("localhost:8080", nil)
	if nil != err {
		panic(err)
	}
}

// helloWorld handler
func helloWorldHandler(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprint(w, "hello world")
	if err != nil {
		fmt.Printf("HelloWorld Server Response Err->%v\n", err)
	}
}

// 复杂andler
func complexHandler(w http.ResponseWriter, req *http.Request) {

	//1.设置响应头
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//2.获取请求
	method := req.Method

	//3.按照不同的方法进行处理
	switch method {
	case http.MethodGet:
		fmt.Fprint(w, formData)
		break
	case http.MethodPost:
		if data, err := io.ReadAll(req.Body); err == nil {
			id := gojsonq.New().FromString(string(data)).Find("id")
			if id == float64(500) {
				panic("server stop")
			}
			fmt.Fprint(w, id)
		}
		break
	default:
		fmt.Printf("method %v not found\n", method)
	}
}

// 初始化handler
func initHandler(f func(w http.ResponseWriter, req *http.Request)) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		//1.defer处理异常
		defer func() {
			if err := recover(); err != nil {
				log.Printf("catch error->[%v]", err)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		//2.执行参数方法
		f(writer, request)
	}
}
