// @Author:冯铁城 [17615007230@163.com] 2023-06-08 09:20:09
package _goroutine

import (
	"fmt"
	"sync"
	"time"
)

// 测试结构体
type info struct {
	a    int
	lock sync.Mutex
}

func (i *info) addA() {
	i.lock.Lock()
	i.a++
	i.lock.Unlock()
}

// TestSync 测试同步锁
func TestSync() {

	//1.创建info对象
	info := new(info)

	//2.循环创建1000个协程 ++a
	for i := 0; i < 1000; i++ {
		go info.addA()
	}

	//3.等待1s，确保所有协程执行成功
	time.Sleep(time.Second)

	//4.输出a
	fmt.Println(info.a)
}
