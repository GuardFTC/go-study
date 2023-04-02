// Package _interface @Author:冯铁城 [17615007230@163.com] 2023-04-02 10:46:39
package _interface

import (
	"fmt"
)

// 定义空接口
type any interface{}

// EmptyInterface 空接口调用
func EmptyInterface() {

	//1.定义一个空接口类型
	var _any any

	//2.指向int
	_any = 1
	fmt.Printf("the _any value is %v\n", _any)

	//2.指向string
	_any = "ftc"
	fmt.Printf("the _any value is %v\n", _any)

	//3.指向结构体
	_any = new(tiger)
	fmt.Printf("the _any value is %v\n", _any)

	//4.类型判定
	switch _any.(type) {
	case int:
		fmt.Printf("the _any type is int\n")
		break
	case string:
		fmt.Printf("the _any type is string\n")
		break
	case *tiger:
		fmt.Printf("the _any type is *tiger\n")
		break
	default:
		fmt.Printf("unknown type\n")
		break
	}

	//5.创建任包含不同类型变量的切片
	anies := make([]any, 0)
	anies = append(anies, 1)
	anies = append(anies, "ftc")
	anies = append(anies, &tiger{})
	fmt.Printf("the anies is %v\n", anies)

	//6.接口到接口，只要底层类型实现了对应的方法，否则会出现运行时错误，注意是运行时
	m := _any.(mammal)
	m.speak()
	m.childbirth()
}
