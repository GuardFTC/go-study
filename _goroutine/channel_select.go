// Package _goroutine @Author:冯铁城 [17615007230@163.com] 2023-05-18 15:40:54
package _goroutine

import (
	"fmt"
	"time"
)

// TestChannelSelect select阻塞通道
func TestChannelSelect() {

	//1.初始化channel
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	//2.创建协程向通道推送数据
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- 1
	}()
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- 2
	}()
	go func() {
		time.Sleep(300 * time.Millisecond)
		ch3 <- 3
	}()

	//3.使用select阻塞式等待通道响应
	flag := false
	for {
		select {
		case <-ch1:
			fmt.Printf("收到第1个通道的消息\n")
		case <-ch2:
			fmt.Printf("收到第2个通道的消息\n")
		case <-ch3:
			fmt.Printf("收到第3个通道的消息\n")
			flag = true
			break
		}

		if flag {
			break
		}
	}
}

// TestChannelSelectV2 select阻塞通道
func TestChannelSelectV2() {

	//1.创建channel
	ch := make(chan string)
	defer func() {
		close(ch)
	}()

	//2.创建第一个人
	go person1(ch)

	//3.创建第二个人
	go person2(ch)

	//4.睡1s
	time.Sleep(1 * time.Second)
}

func person1(ch chan string) {

	//1.发送第一句话
	ch <- "你好"

	//2.监听队列，回复消息
	for {
		select {
		case value := <-ch:
			fmt.Printf("person2 sad:%v\n", value)
			switch value {
			case "你好":
				ch <- "你叫什么名字"
			case "王钢蛋":
				ch <- "你好我叫赵铁柱"
			case "明天见":
				ch <- "好滴明天见"
				goto L
			}
		}
	}
L:
}

func person2(ch chan string) {
	for {
		select {
		case value := <-ch:
			fmt.Printf("person1 sad:%v\n", value)
			switch value {
			case "你好":
				ch <- "你好"
			case "你叫什么名字":
				ch <- "王钢蛋"
			case "你好我叫赵铁柱":
				ch <- "明天见"
			case "好滴明天见":
				goto L
			}
		}
	}
L:
}
