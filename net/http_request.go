// Package net @Author:冯铁城 [17615007230@163.com] 2023-06-20 09:13:47
package net

import (
	"fmt"
	"io"
	"net/http"

	"github.com/thedevsaddam/gojsonq/v2"
)

// ecosys服务器响应结构体
type ecosysResponse struct {
	code int
	msg  string
	data any
}

// TestGet 测试Get请求
func TestGet() {

	//定义URL
	url := "https://ecosys.cicic.cn/api/wechat/open-ids?code=021bAdHa1p62iF0VwVGa1FJg300bAdHu"

	//发送GET请求
	resp, err := http.Get(url)

	if nil != err {
		fmt.Printf("request err->%v\n", err)
		return
	}

	//获取响应码
	statusCode := resp.StatusCode

	//获取响应体
	data, err := io.ReadAll(resp.Body)
	if nil != err {
		fmt.Printf("parse body err->%v\n", err)
		return
	}
	bodyStr := string(data)

	//打印响应码以及响应体
	fmt.Printf("respStatus is %v, respBody is %v\n", statusCode, bodyStr)

	//将响应体转换为json
	fmt.Printf("msg is %v\n", gojsonq.New().FromString(bodyStr).Find("msg"))
	fmt.Printf("code is %v\n", gojsonq.New().FromString(bodyStr).Find("code"))
	fmt.Printf("body is %v\n", gojsonq.New().FromString(bodyStr).Find("data"))
}
