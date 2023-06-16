// Package _goroutine @Author:冯铁城 [17615007230@163.com] 2023-05-19 15:54:59
package _goroutine

import (
	"fmt"
)

// 定义通道
var chanFibonacci chan int64

// 定义入队数
var inChanNumber int64 = 1

// 定义暂存数
var storageNumber int64 = 0

// TestLazyGenerator 测试惰性生成器
func TestLazyGenerator() {

	//1.生成菲波那切channel
	chanFibonacci = fibonacci()

	//2.循环取值
	for i := 0; i < 20; i++ {
		fmt.Printf("%v, ", getDataFromFibonacci())
	}
}

// 获取斐波那契数通道
// 解题思路
// 入队数  暂存数
// 1        0
// 1+0=1    1
// 1+1=2    1
// 2+1=3    2
func fibonacci() chan int64 {

	//1.创建通道
	channel := make(chan int64)

	//2.创建协程存入初始数
	go func() {
		for {

			//3.存入通道
			channel <- inChanNumber

			//4.计算新的入队数
			out := inChanNumber
			inChanNumber = out + storageNumber

			//5.变更暂存数
			storageNumber = out
		}
	}()

	//6.返回通道
	return channel
}

// 从斐波那契通道取值
func getDataFromFibonacci() int64 {
	return <-chanFibonacci
}
