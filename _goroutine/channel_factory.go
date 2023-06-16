// Package _goroutine @Author:冯铁城 [17615007230@163.com] 2023-05-18 15:40:54
package _goroutine

import (
	"fmt"
)

// TestChanFactory 测试通道工厂模式
func TestChanFactory() {

	//1.获取发送数据通道
	chanSendData := testChanFactorySendData()

	//2.获取读取数据通道
	chanGetData := testChanFactoryGetData(chanSendData)

	//3.阻塞等待信号
	<-chanGetData
}

func testChanFactorySendData() <-chan int {

	//1.创建通道
	ch := make(chan int)

	//2.创建协程向通道推送10个数据
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i * 5
		}
		close(ch)
	}()

	//3.返回通道
	return ch
}

func testChanFactoryGetData(ch <-chan int) chan bool {

	//1.创建通道
	chBool := make(chan bool)

	//2.创建协程读取参数通道数据，读取完成后向信号量通道推送数据
	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Printf("channel data is %v\n", <-ch)
		}
		chBool <- true
		close(chBool)
	}()

	//3.返回通道
	return chBool
}
