// Package datatype @Author:冯铁城 [17615007230@163.com] 2023-03-03 10:16:01
package datatype

import "fmt"

// 声明全局变量
var id = "ftc"

// 批量声明全局变量
var (
	game   = "LOL"
	player = "faker"
)

// 声明未赋值变量,通过init函数赋值
var value string

// init函数
func init() {
	value = "initValue"
}

// PrintVariable 打印变量函数
func PrintVariable() {
	fmt.Println("------------------------Console Variable---------------------------")

	// 打印变量
	fmt.Println(id)
	fmt.Println(game)
	fmt.Println(player)
	fmt.Println(value)

	// 定义函数内部局部变量
	game := "NBA2K"
	fmt.Println(game)

	// 修改player变量的值
	player = "curry"
	fmt.Println(player)

	// 定义局部变量a,b
	a, b, c := 1, 2, 3
	fmt.Printf("a=%v,b=%v,c=%v\n", a, b, c)

	// 交换ab的值
	a, b, c = c, a, b
	fmt.Printf("a=%v,b=%v,c=%v\n", a, b, c)

	//输出a,b,c的内存地址
	fmt.Printf("location(a)=%v,location(b)=%v,location(c)=%v\n", &a, &b, &c)
}
