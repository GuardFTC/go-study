// Package _goroutine @Author:冯铁城 [17615007230@163.com] 2023-05-18 15:40:54
package _goroutine

import (
	"fmt"
	"time"
)

// TestChannelAsync 非阻塞式channel
func TestChannelAsync() {

	//1.声明一个channel
	var ch chan int

	//2.初始化一个有缓冲区，大小为5，非阻塞式channel
	ch = make(chan int, 5)

	//3.创建发送数据协程
	go sendDataAsync(ch)

	//4.创建接收数据协程
	go getDataAsync(ch)

	//5.睡3s等待协程执行完成
	time.Sleep(3 * time.Second)
}

// 非阻塞式发送数据
func sendDataAsync(ch chan int) {
	start := time.Now().Unix()
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("send data=%v async\n", i)
	}
	end := time.Now().Unix()
	fmt.Printf("send data async success. cast time=%vs\n", end-start)
}

// 非阻塞式接收数据
func getDataAsync(ch chan int) {
	time.Sleep(2 * time.Second)
	for i := 0; i < 5; i++ {
		fmt.Printf("get data=%v async\n", <-ch)
	}
}
