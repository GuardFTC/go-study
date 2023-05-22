// Package _goroutine @Author:冯铁城 [17615007230@163.com] 2023-05-22 15:41:04
package _goroutine

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"time"
)

// TestChain 测试链式通道
func TestChain() {

	//1.开始计时
	start := time.Now().UnixMilli()

	//2.定义循环次数
	maxNumber := 100000

	//3.定义通道
	rightmost := make(chan int)

	//4.初始化指针
	var left, right chan int = rightmost, nil

	//5.循环操作
	for i := 0; i < maxNumber; i++ {

		//6.声明left
		right, left = left, make(chan int)

		//7.创建新协程操作
		//PS:此处如果用匿名函数(闭包)，会报错all goroutines are asleep - deadlock!
		go mulAndPushData(left, right)
	}

	//8.从入口处推入一个数
	left <- 0

	//9.阻塞式等待从最右队列获取结果值
	result := <-rightmost
	fmt.Printf("result is %v\n", result)

	//10.计算总耗时
	end := time.Now().UnixMilli()
	fmt.Printf("cost time is %vms\n", end-start)
}

// 从左侧通道取值，存入右侧通道
func mulAndPushData(left, right chan int) {
	right <- 1 + <-left
}

// 获取协程ID
func getCoroutineId() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
