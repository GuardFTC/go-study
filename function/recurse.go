// Package function @Author:冯铁城 [17615007230@163.com] 2023-03-13 17:58:03
package function

import "fmt"

// Recurse 递归打印99乘法表
func Recurse(startNum int) {

	//1.递归结束条件
	if startNum > 9 {
		return
	}

	//2.打印当前数值
	for i := startNum; i <= 9; i++ {
		fmt.Printf("%d*%d=%d ", startNum, i, startNum*i)
	}
	fmt.Println()

	//3.递归
	Recurse(startNum + 1)
}
