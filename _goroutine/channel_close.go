// Package _goroutine @Author:冯铁城 [17615007230@163.com] 2023-05-18 15:40:54
package _goroutine

import (
	"fmt"
)

// TestChannelClose 关闭通道
func TestChannelClose() {

	//1.初始化一个channel
	ch := make(chan int)

	//2.创建一个协程发送数据
	go sendDataAndClose(ch)

	//3.主协程接收数据
	getDataAndClose(ch)
}

func sendDataAndClose(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func getDataAndClose(ch chan int) {
	for {
		value, open := <-ch
		if !open {
			break
		}
		fmt.Printf("get data value=%v\n", value)
	}
}
