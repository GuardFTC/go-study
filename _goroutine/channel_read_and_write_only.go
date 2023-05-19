// Package _goroutine @Author:冯铁城 [17615007230@163.com] 2023-05-18 15:40:54
package _goroutine

import "fmt"

// TestReadAndWriteOnly 测试只读/只写通道模式
func TestReadAndWriteOnly() {

	//1.创建通道
	channel := make(chan int)

	//2.向通道写入数据
	go writeOnlyData(channel)

	//3.从通道读取数据
	chanSignal := readOnlyData(channel)

	//4.监听信号，结束协程
	<-chanSignal
}

// 向只写通道写入数据
func writeOnlyData(writeOnlyChan chan<- int) {
	for i := 0; i < 10; i++ {
		writeOnlyChan <- i
		//i := <-writeOnlyData 只写通道无法读取数据
	}
}

// 从只读通道读取数据
func readOnlyData(readOnlyChan <-chan int) chan bool {

	//1.创建信号量通道
	signal := make(chan bool)

	//2.创建新协程从只读通道读取数据
	go func() {
		readCount := 0
		for v := range readOnlyChan {
			fmt.Printf("read data from read_only_chan=%v\n", v)
			if readCount == 9 {
				signal <- true
				break
			}
			readCount++
		}
	}()
	//readOnlyChan <- 1 只读通道无法写入数据

	//3.信号量通道
	return signal
}
