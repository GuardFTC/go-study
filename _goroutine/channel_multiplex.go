// Package _goroutine @Author:冯铁城 [17615007230@163.com] 2023-05-22 11:01:45
package _goroutine

import (
	"fmt"
	"time"
)

// Request 请求类型结构体
type Request struct {
	a, b    int
	reqChan chan int
}

// Server 服务器结构体
type Server struct {
	doFunc func(int, int) int
}

// 定义服务器开启方法
func (s *Server) start() (chan *Request, chan bool) {

	//1.创建通道
	requestsChan := make(chan *Request)
	serverStatusChan := make(chan bool)

	//2.创建新协程监听通道信息
	go func() {
		for {
			select {
			case req := <-requestsChan:
				go s.dealRequest(req)
			case <-serverStatusChan:
				s.close()
				return
			}
		}
	}()

	//3.返回通道
	return requestsChan, serverStatusChan
}

// 定义服务器关闭方法
func (s *Server) close() {
	fmt.Println("server is closed")
}

// 定义处理请求方法
func (s *Server) dealRequest(req *Request) {
	req.reqChan <- s.doFunc(req.a, req.b)
}

// TestMultiplex 测试多路复用
func TestMultiplex() {

	//1.定义服务器处理逻辑
	mulFunc := func(a, b int) int {
		return a + b
	}

	//2.创建服务器指针类型变量
	server := new(Server)
	server.doFunc = mulFunc

	//3.开启服务器
	requestsChan, serverStatusChan := server.start()

	//4.定义请求切片
	requestNumber := 100
	requests := make([]*Request, requestNumber)

	//5.推入100个请求
	for i := 0; i < requestNumber; i++ {

		//6.创建请求
		request := new(Request)
		request.a = i
		request.b = i
		request.reqChan = make(chan int)

		//7.存入切片
		requests[i] = request

		//8.发送给服务器
		requestsChan <- request
	}

	//9.验证结果
	for i := 0; i < requestNumber; i++ {
		if i+i == <-requests[i].reqChan {
			fmt.Printf("the request is ok. result=%v\n", i+i)
		} else {
			fmt.Printf("the request is error. index is %v", i)
		}
	}

	//10.关闭服务器
	serverStatusChan <- true

	//11.睡100ms，等待服务器协程结束
	time.Sleep(100 * time.Millisecond)
}
