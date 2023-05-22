// Package _goroutine @Author:冯铁城 [17615007230@163.com] 2023-05-18 15:40:54
package _goroutine

import (
	"fmt"
	"time"
)

// TestFuture 测试future
func TestFuture() {

	//1.定义参数
	param := 4

	//2.开启计时
	start := time.Now().UnixMilli()

	//3.异步任务乘法
	mulChannel := testFutureMul(param)

	//4.异步任务除法
	divChannel := testFutureDiv(param)

	//5.睡1s
	time.Sleep(time.Second)

	//6.阻塞式获取结果
	mul := <-mulChannel
	div := <-divChannel

	//7.输出最终结果
	fmt.Printf("the final result:%v\n", mul+div)

	//8.打印最终耗时
	end := time.Now().UnixMilli()
	fmt.Printf("this func cost %vms", end-start)
}

// 乘法异步任务
func testFutureMul(param int) chan int {

	//1.创建channel
	intChan := make(chan int)

	//2.创建新的协程将param*2
	go func() {

		//3.睡1s，模拟操作耗时
		time.Sleep(time.Second)

		//4.将结果推入通道
		intChan <- param << 1
	}()

	//5.返回通道
	return intChan
}

// 除法异步任务
func testFutureDiv(param int) chan int {

	//1.创建channel
	intChan := make(chan int)

	//2.创建新的协程将param/2
	go func() {

		//3.睡2s，模拟操作耗时
		time.Sleep(2 * time.Second)

		//4.将结果推入通道
		intChan <- param >> 1
	}()

	//5.返回通道
	return intChan
}
