// Package function @Author:冯铁城 [17615007230@163.com] 2023-03-13 18:18:34
package function

import (
	"fmt"
	"log"
	"runtime"
)

// TestClosure 测试闭包的使用
func TestClosure() {

	//1.直接调用匿名函数
	max := func(a int, b int) (max int) {
		if a > b {
			max = a
		} else {
			max = b
		}
		return max
	}(2, 3)
	fmt.Printf("max=%d\n", max)

	//2.定义匿名函数并赋值给变量，通过变量调用
	f := func(a int, b int) (max int) {
		if a > b {
			max = a
		} else {
			max = b
		}
		return max
	}
	fmt.Printf("f'type is %T\nf(2,3)=%d\n", f, f(2, 3))

	//3.应用闭包-函数作为返回值
	addTwo := addTwo(3)
	fmt.Printf("addTwo's type is %T\naddTwo(2) = %d\n", addTwo, addTwo(2))

	//4.应用闭包-跟踪函数进程
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s", file, line)
	}
	where()
	where()
}

// 与任意数字+
func addTwo(a int) func(b int) (sum int) {
	return func(b int) (sum int) {
		return a + b
	}
}
