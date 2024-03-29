// Package _goroutine @Author:冯铁城 [17615007230@163.com] 2023-05-18 15:40:54
package _goroutine

import (
	"fmt"
	"time"
)

// TestChannelTimer 非阻塞式channel
func TestChannelTimer() {

	//1.创建周期性秒表
	tick := time.Tick(1 * time.Second)

	//2.创建一次性秒表
	after := time.After(5 * time.Second)

	//3.select阻塞监听
	for true {
		select {
		case <-tick:
			fmt.Printf("now is %v\n", time.Now().Format(time.DateTime))
		case <-after:
			fmt.Println("is over!!!!!")
			goto L
		}
	}
L:
}

// TestTimeout 测试超时监听
func TestTimeout() {

	//1.初始化一个channel
	channel := make(chan bool, 1)
	defer close(channel)

	//2.创建一次性秒表，延迟3s
	after := time.After(3 * time.Second)

	//3.模拟从数据库读取数据，阻塞时间为5s
	go getDataFromDatabase(channel)

	//4.开始监听
	for {
		select {
		case v := <-channel:
			fmt.Printf("get channel data=%v\n", v)
			goto L
		case <-after:
			fmt.Printf("get data timeout\n")
			goto L
		}
	}
L:
}

// 模拟从数据库读取数据
func getDataFromDatabase(ch chan bool) {

	//1.睡5s
	time.Sleep(5 * time.Second)

	//2.向通道发送数据
	ch <- true
}
