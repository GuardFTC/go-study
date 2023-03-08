// Package datatype @Author:冯铁城 [17615007230@163.com] 2023-03-03 09:38:22
package datatype

import "fmt"

// Name 定义常量
const name = "ftc"

// Age 定义常量计算
const age = 10 / 2

// 定义枚举类型常量
const (
	ok    = 1
	notOk = 0
)

// PrintConstant 打印常量函数
func PrintConstant() {
	fmt.Println("------------------------Console Constant---------------------------")
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(ok)
	fmt.Println(notOk)
}
