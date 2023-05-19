// Package _goroutine @Author:冯铁城 [17615007230@163.com] 2023-05-18 15:12:22
package _goroutine

import (
	"fmt"
	"time"
)

// TryGoroutine 协程测试代码
func TryGoroutine() {
	fmt.Println("tryGoroutine sleep start")

	//1.开启协程异步执行
	go longSleep()
	go shortSleep()

	//2.执行自己的内容
	time.Sleep(10 * 1e9)
	fmt.Println("tryGoroutine sleep end")

}

// 长期休眠
func longSleep() {
	fmt.Println("longSleep start")
	time.Sleep(5 * 1e9)
	fmt.Println("longSleep end")
}

// 短期休眠
func shortSleep() {
	fmt.Println("shortSleep start")
	time.Sleep(3 * 1e9)
	fmt.Println("shortSleep end")
}
