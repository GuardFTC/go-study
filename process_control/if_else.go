// Package process_control @Author:冯铁城 [17615007230@163.com] 2023-03-08 10:26:28
package process_control

import "fmt"

// IfElse if-else流程控制的各种写法
func IfElse() {

	//1.最简单的流程控制
	a := 1
	if a == 1 {
		fmt.Printf("when a'value is %d,run the if process\n", a)
	} else if a == 2 {
		fmt.Printf("when a'value is %d,run the else if process\n", a)
	} else {
		fmt.Printf("when a'value is %d,run the else process\n", a)
	}

	//2.缩略写法
	if a := 2; a == 1 {
		fmt.Printf("when a'value is %d,run the if process\n", a)
	} else if a == 2 {
		fmt.Printf("when a'value is %d,run the else if process\n", a)
	} else {
		fmt.Printf("when a'value is %d,run the else process\n", a)
	}

	//3.异常判定写法
	if _, err := fmt.Println(1); err != nil {
		fmt.Println("when error is not nil,run the if process")
	} else {
		fmt.Println("when error is nil,run the else process")
	}

	//4.最常见的ok写法
	if _, ok := test(3); ok {
		fmt.Println("when ok is true,run the if process")
	} else {
		fmt.Println("when ok is false,run the else process")
	}
}

// 测试方法
func test(i int) (int, bool) {
	if i > 0 {
		return 1, true
	} else {
		return -1, false
	}
}
