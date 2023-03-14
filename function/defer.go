// Package function @Author:冯铁城 [17615007230@163.com] 2023-03-13 12:20:03
package function

import "fmt"

func TestDefer() {

	//1.基础调用
	fmt.Printf("deferTest param:[a=%d,b=%d] return value:[%d]\n", 6, 7, deferTest(6, 7))

	//2.循环调用
	deferTest2()
}

// 基础调用
func deferTest(a int, b int) (max int) {
	defer func() {
		fmt.Printf("max value is %d\n", max)
	}()
	if a > b {
		max = a
	} else {
		max = b
	}
	return max
}

// 循环defer
func deferTest2() {
	defer fmt.Println()
	for i := 0; i < 6; i++ {
		defer fmt.Printf("%d ", i)
	}
}
