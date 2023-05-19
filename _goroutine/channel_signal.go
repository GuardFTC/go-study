// Package _goroutine @Author:冯铁城 [17615007230@163.com] 2023-05-18 15:40:54
package _goroutine

import (
	"fmt"
	"time"
)

// TestSignal 测试信号量模式
func TestSignal() {

	//1.创建一个无缓冲区阻塞式channel
	ch := make(chan bool)

	//2.创建协程执行异步逻辑
	go testSignalSend(ch)

	//3.获取当前时间戳
	start := time.Now().Unix()

	//4.打印日志
	fmt.Println("等待异步任务执行完成")

	//4.作为消费者阻塞式读取通道数据,等待信号
	<-ch

	//5.获取结束时间
	end := time.Now().Unix()

	//6.打印日志
	fmt.Printf("异步任务执行完成 总耗时:%vs", end-start)
}

// 测试信号量模式-发送数据
func testSignalSend(ch chan bool) {

	//1.睡2s
	time.Sleep(2 * time.Second)

	//2.向通道发送信号
	ch <- true
}
