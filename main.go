// @Author:冯铁城 [17615007230@163.com] 2023-03-08 18:24:43
package main

import _interface "go-study/interface"

func main() {

	//变量类型学习
	dataTypeStudy()

	//流程控制学习
	processControlStudy()

	//方法学习
	functionStudy()

	//切片学习
	sliceStudy()

	//map学习
	mapStudy()

	//结构体学习
	structStudy()

	//接口学习
	interfaceStudy()
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

	//获取函数类型
	//function.GetFuncType()

	//简单函数应用
	//function.TestMax()

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

// 切片学习
func sliceStudy() {

	//array学习
	//slice.CreatArray()

	//array内存分析
	//slice.MemoryArray()

	//slice简单学习
	//slice.CreateSlice()

	//slice的Append操作
	//slice.UseAppend()

	//slice的内存分析
	//slice.MemorySlice()

	//slice的扩容机制
	//slice.EnlargeSlice()
}

// map学习
func mapStudy() {

	//创建map
	//map_study.CreateMap()
}

// 结构体学习
func structStudy() {

	//创建结构体
	//_struct.CreateStruct()

	//结构体内存分析
	//_struct.StructMemory()

	//结构体标签操作
	//_struct.StructTag()

	//内嵌结构体
	//_struct.AmbiguousStruct()

	//结构体方法
	//_struct.TestFunction()
}

// 接口学习
func interfaceStudy() {

	//接口的基础使用
	//_interface.UseInterface()

	//接口类型断言
	//_interface.InterfaceAssertCheck()

	//空接口
	//_interface.EmptyInterface()

	//反射
	//_interface.Ref()

	//结构体反射
	_interface.StructRef()

	//自定义场景测试
	//test.Gaming()
}
