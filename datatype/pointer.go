// Package datatype @Author:冯铁城 [17615007230@163.com] 2023-03-05 16:53:14
package datatype

import "fmt"

// PrintPointer 指针相关类型输出
func PrintPointer() {

	//创建一个int类型变量
	number := 1

	//输出number的内存地址
	fmt.Printf("the location of number(%d) in memory is %v\n", number, &number)

	//创建一个int类型指针
	var ptr *int

	//输出指针的内存地址
	fmt.Printf("the location of ptr(%v) in memory is %v\n", ptr, &ptr)

	//输出指针的值，因为指针为空，因此会报错：panic: runtime error: invalid memory address or nil pointer dereference
	//fmt.Printf("ptr's value is %v\n", *ptr)

	//ptr指向number
	ptr = &number
	fmt.Printf("the location of ptr(%v) in memory is %v\n", ptr, &ptr)

	//输出指针的值
	fmt.Printf("ptr's value is %v\n", *ptr)
}
