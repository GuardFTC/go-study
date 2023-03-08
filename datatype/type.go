// Package datatype @Author:冯铁城 [17615007230@163.com] 2023-03-03 20:35:58
package datatype

import (
	"fmt"
	"math/rand"
	"unicode"
)

// 定义布尔类型变量
var isTrue = true
var isFalse = false

// 定义整型变量
var intA int8 = 127
var intB int16 = 32767
var intC int32 = 2147483647
var intD int64 = 922337036854775807

// 定义浮点型变量
var floatA float32 = 2.2222222
var floatB = 3.333333333333333

// 定义复数
var c1 complex64 = 2 + 10i
var c2 complex128 = 30 + 100i

// FTC 定义类型别名=int64
type FTC int64

// 定义FTC类型变量
var ftcA FTC = 66666

// 定义字符类型
var c = 'A'

// PrintBoolean 输出布尔类型
func PrintBoolean() {
	fmt.Println("------------------------Console Boolean---------------------------")
	fmt.Printf("isTrue = %t isFalse = %t \n", !isTrue, isFalse)

	// 打印boolean值
	fmt.Printf("!isTrue == %t\n", !isTrue)
	fmt.Printf("!isFalse == %t\n", !isFalse)

	// 打印boolean值
	fmt.Printf("true && true == %t\n", isTrue && isTrue)
	fmt.Printf("false && false == %t\n", isFalse && isFalse)
	fmt.Printf("false && true == %t\n", isFalse && isTrue)
	fmt.Printf("true && false == %t\n", isTrue && isFalse)

	// 打印boolean值
	fmt.Printf("true || true == %t\n", isTrue || isTrue)
	fmt.Printf("false || false == %t\n", isFalse || isFalse)
	fmt.Printf("false || true == %t\n", isFalse || isTrue)
	fmt.Printf("true || false == %t\n", isTrue || isFalse)
}

// PrintIntAndFloat 打印整型和浮点型
func PrintIntAndFloat() {
	fmt.Println("------------------------Console Int And Float---------------------------")

	// 输出整型以及浮点型
	fmt.Printf("intA(bit8) == %d\n", intA)
	fmt.Printf("intB(bit16) == %d\n", intB)
	fmt.Printf("intC(bit32) == %d\n", intC)
	fmt.Printf("intD(bit64) == %0d\n", intD)
	fmt.Printf("intD(bit64) == %50d\n", intD)
	fmt.Printf("floatA(bit32) == %g\n", floatA)
	fmt.Printf("floatB(bit64) == %g\n", floatB)

	// 定义临时浮点型数字,并格式化输出（控制小数点数）
	a := 3.33333333333
	fmt.Printf("format a == %0.4g\n", a)
	fmt.Printf("format a == %50.4g\n", a)

	// 创建局部变量，进行类型转换，计算后输出
	b := int16(intA)
	fmt.Printf("intB-b == %d\n", intB-b)
}

// PrintComplex 打印复数
func PrintComplex() {
	fmt.Println("------------------------Console Complex---------------------------")
	fmt.Printf("c1 = %v\n", c1)
	fmt.Printf("c2 = %v\n", c2)
}

// PrintOperation 打印运算符
func PrintOperation() {
	fmt.Println("------------------------Console Operation---------------------------")

	// 按位与
	fmt.Printf("1 & 1 == %d\n", 1&1)
	fmt.Printf("0 & 0 == %d\n", 0&0)
	fmt.Printf("1 & 0 == %d\n", 1&0)
	fmt.Printf("0 & 1 == %d\n", 0&1)

	// 按位或
	fmt.Printf("1 | 1 == %d\n", 1|1)
	fmt.Printf("0 | 0 == %d\n", 0|0)
	fmt.Printf("1 | 0 == %d\n", 1|0)
	fmt.Printf("0 | 1 == %d\n", 0|1)

	// 按位异或
	fmt.Printf("1 ^ 1 == %d\n", 1^1)
	fmt.Printf("0 ^ 0 == %d\n", 0^0)
	fmt.Printf("1 ^ 0 == %d\n", 1^0)
	fmt.Printf("0 ^ 1 == %d\n", 0^1)

	// 位移运算
	fmt.Printf("1<<2 == %d\n", 1<<2)
	fmt.Printf("4>>1 == %d\n", 4>>1)
}

// PrintRandom 输出随机数
func PrintRandom() {
	fmt.Println("------------------------Console RandomIn---------------------------")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", rand.Intn(10000))
	}
	fmt.Println()

	fmt.Println("------------------------Console RandomFloat---------------------------")
	for i := 0; i < 5; i++ {
		fmt.Printf("%g ", rand.Float32())
	}
	fmt.Println()
}

// PrintChar 打印字符类型
func PrintChar() {
	fmt.Println("------------------------Console Char---------------------------")
	fmt.Printf("c == %c\n", c)
	fmt.Printf("'2' is digit?:%t\n", unicode.IsDigit('2'))
	fmt.Printf("'A' is letter?:%t\n", unicode.IsLetter('A'))
	fmt.Printf("' ' is space?:%t\n", unicode.IsSpace(' '))
	fmt.Printf("'A' is lower?:%t\n", unicode.IsLower('A'))
	fmt.Printf("'a' is upper?:%t\n", unicode.IsUpper('a'))
}

// PrintTypeAlias 打印类型别名
func PrintTypeAlias() {
	fmt.Println("------------------------Console TypeAlias---------------------------")
	fmt.Printf("%d\n", ftcA)
}
