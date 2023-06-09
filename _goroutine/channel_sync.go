// Package _goroutine @Author:冯铁城 [17615007230@163.com] 2023-05-18 15:40:54
package _goroutine

import (
	"fmt"
	"time"
)

// TestChannelSync  阻塞式channel
func TestChannelSync() {

	//1.声明一个channel
	var ch chan int

	//2.初始化一个无缓冲区，阻塞式channel
	ch = make(chan int)

	//3.创建发送数据协程
	go sendDataSync(ch)

	//4.创建接收数据协程
	go getDataSync(ch)

	//5.睡3s等待协程执行完成
	time.Sleep(3 * time.Second)
}

// 阻塞式发送数据
func sendDataSync(ch chan int) {
	start := time.Now().Unix()
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("send data=%v sync\n", i)
	}
	end := time.Now().Unix()
	fmt.Printf("send data sync success. cast time=%vs\n", end-start)
}

// 阻塞式接收数据
func getDataSync(ch chan int) {
	time.Sleep(2 * time.Second)
	for i := 0; i < 5; i++ {
		fmt.Printf("get data=%v sync\n", <-ch)
	}
}
