// Package _goroutine @Author:冯铁城 [17615007230@163.com] 2023-06-19 10:32:10
package _goroutine

import (
	"fmt"
	"runtime"
	"time"
)

// TestMemoryLeak 测试内存溢出
func TestMemoryLeak(f func()) {

	//1.打印最终协程数
	defer func() {
		fmt.Printf("the count of goroutine is %v\n", runtime.NumGoroutine())
	}()

	//2.打印初始协程数
	fmt.Printf("the count of goroutine is %v\n", runtime.NumGoroutine())

	//3.执行参数方法
	f()

	//4.睡3s
	time.Sleep(3 * time.Second)
}

// TestMemoryLeakOneSenderOrSendOneTimes 模拟内存溢出 1个发送者或发送一次
func TestMemoryLeakOneSenderOrSendOneTimes() {

	//1.创建通道
	ch := make(chan int)
	//defer close(ch)

	//2.创建新协程写入数据，阻塞2s
	go func() {
		//defer close(ch)
		time.Sleep(2 * time.Second)
		ch <- 1
		fmt.Println("send data success")
	}()

	//3.监听ch，超时时间为1s
	select {
	case i := <-ch:
		fmt.Printf("receive data is %v\n", i)
	case <-time.After(time.Second):
		fmt.Println("receive data timeout")
	}
}

// TestMemoryLeakOneSenderOrSendOneTimesOk 模拟内存溢出 1个发送者或发送一次
func TestMemoryLeakOneSenderOrSendOneTimesOk() {

	//1.创建通道
	ch := make(chan int, 1)

	//2.创建新协程写入数据，阻塞2s
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
		fmt.Println("send data success")
	}()

	//3.监听ch，超时时间为1s
	select {
	case i := <-ch:
		fmt.Printf("receive data is %v\n", i)
	case <-time.After(time.Second):
		fmt.Println("receive data timeout")
	}
}

// TestMemoryLeakOneListenerOrListenOneTimes 模拟内存溢出 1个监听者或监听一次
func TestMemoryLeakOneListenerOrListenOneTimes() {

	//1.创建通道
	ch := make(chan int)

	//2.创建新协程监听数据，阻塞2s
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Printf("receive data is %v\n", <-ch)
		fmt.Println("receive data end")
	}()

	//3.关闭通道
	close(ch)
}

// TestMemoryLeakMoreSenderOrSendMoreTimes 模拟内存溢出 多个发送者或发送多次
func TestMemoryLeakMoreSenderOrSendMoreTimes() {

	//1.创建通道
	ch := make(chan int, 1)

	//2.创建新协程写入数据
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- i
		}
		fmt.Println("send data success")
	}()

	//3.监听通道，读取数据
	for value := range ch {
		fmt.Printf("receive data is %v\n", value)
		if value == 2 {
			break
		}
	}
	fmt.Println("receive data end")
}

// TestMemoryLeakMoreSenderOrSendMoreTimesOk 模拟内存溢出 多个发送者或发送多次
func TestMemoryLeakMoreSenderOrSendMoreTimesOk() {

	//1.创建通道
	ch := make(chan int, 1)
	statusChan := make(chan bool)
	defer close(statusChan)

	//2.创建新协程写入数据
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			select {
			case <-statusChan:
				goto L
			default:
				ch <- i
			}
		}
	L:
		fmt.Println("send data success")
	}()

	//3.监听通道，读取数据
	for value := range ch {
		fmt.Printf("receive data is %v\n", value)
		if value == 2 {
			statusChan <- true
			break
		}
	}
	fmt.Println("receive data end")
}

// TestMemoryLeakMoreListenerOrListenMoreTimes 模拟内存溢出 多个监听者或监听多次
func TestMemoryLeakMoreListenerOrListenMoreTimes() {

	//1.创建通道
	ch := make(chan int, 1)
	defer close(ch)

	//2.创建新协程监听数据
	go func() {
		for value := range ch {
			fmt.Printf("receive data is %v\n", value)
		}
		fmt.Println("receive data end")
	}()

	//3.写入数据
	for i := 0; i < 5; i++ {
		ch <- i
		if i == 2 {
			break
		}
	}
	fmt.Println("send data success")
}
