// @Author:冯铁城 [17615007230@163.com] 2023-03-08 18:24:43
package main

import "go-study/function"

func main() {

	//变量类型学习
	dataTypeStudy()

	//流程控制学习
	processControlStudy()

	//方法学习
	functionStudy()
}

// 变量类型学习
func dataTypeStudy() {
	//打印常量
	//datatype.PrintConstant()

	//打印变量
	//datatype.PrintVariable()

	//打印boolean类型
	//datatype.PrintBoolean()

	//打印整型以及浮点型
	//datatype.PrintIntAndFloat()

	//打印复数
	//datatype.PrintComplex()

	//打印运算符
	//datatype.PrintOperation()

	//打印随机数
	//datatype.PrintRandom()

	//打印字符类型
	//datatype.PrintChar()

	//打印类型别名
	//datatype.PrintTypeAlias()

	//打印字符串
	//datatype.PrintString()

	//打印时间相关
	//datatype.PrintTime()

	//打印指针相关
	//datatype.PrintPointer()
}

// 流程控制学习
func processControlStudy() {

	//if-else流程控制
	//process_control.IfElse()

	//switch-case流程控制
	//process_control.SwitchCase()

	//for循环流程控制
	//process_control.For()
}

// 方法学习
func functionStudy() {

	//简单函数应用
	function.TestMax()

	//defer应用
	//function.TestDefer()

	//递归应用
	//function.Recurse(1)

	//函数参数应用
	//fmt.Printf("the first digit index in %s is %d\n", "ftc3", strings.IndexFunc("ftc3", unicode.IsDigit))
	//function.TestParamMethod()

	//闭包函数调用
	//function.TestClosure()
}
