// Package _interface @Author:冯铁城 [17615007230@163.com] 2023-04-01 20:02:30
package _interface

import "fmt"

// UseInterface 使用接口
func UseInterface() {

	//1.创建一个animal接口变量
	var a animal

	//2.创建一个tiger结构体变量
	a = &tiger{}

	//3.执行speak方法
	a.speak()

	//4.创建一个切片,用来存储mammal接口变量
	mammals := make([]mammal, 0)

	//5.存入两种实现类
	mammals = append(mammals, new(tiger))
	mammals = append(mammals, new(mouse))

	//6.循环mammals
	for _, m := range mammals {
		m.speak()
		m.childbirth()
	}
}

// InterfaceAssertCheck 接口类型断言检查
func InterfaceAssertCheck() {

	//1.创建一个animal接口类型变量
	var a animal
	a = new(mouse)

	//2.类型断言,验证a的类型,类型验证不一定每次都会成功,所以通过ok表达式的形式来确保准确性
	if t, ok := a.(*tiger); ok {
		fmt.Printf("the a's type is %T\n", t)
	} else {
		fmt.Printf("the a's type is not *tiger\n")
	}

	//3.通过switch-case验证类型,type-switch不允许有fallthrough
	switch t := a.(type) {
	case *tiger:
		fmt.Printf("a's type is tiger:[%T]\n", t)
		break
	case *mouse:
		fmt.Printf("a's type is month:[%T]\n", t)
		break
	default:
		fmt.Printf("unknown type")
		break
	}

	//4.判定变量是否实现了某个接口，需要先将该变量声明为空接口,只有接口类型变量才可以进行类型断言检测
	var obj interface{}
	obj = new(chicken)

	//4.1 判定是否实现了animal接口
	if a, ok := obj.(animal); ok {
		a.speak()
	} else {
		fmt.Printf("obj not implement animal interface")
	}

	//4.2 判定是否实现了mammal接口
	if m, ok := obj.(mammal); ok {
		m.childbirth()
	} else {
		fmt.Printf("obj not implement mammal interface")
	}
}

// 定义动物接口
type animal interface {

	//定义speak方法
	speak()
}

// 定义哺乳动物接口
type mammal interface {

	//定义动物方法
	animal

	//定义分娩方法
	childbirth()
}

// 定义老虎结构体
type tiger struct {
}

// 实现animal接口的speak方法
func (t *tiger) speak() {
	fmt.Println("嗷呜~~~~~~~~~~~~~~~~")
}

// 实现mammal接口的childbirth方法
func (t *tiger) childbirth() {
	fmt.Println("一只小老虎出生了")
}

// 定义老鼠结构体
type mouse struct {
}

// 实现animal接口的speak方法
func (m *mouse) speak() {
	fmt.Println("吱吱！吱吱！吱吱！")
}

// 实现mammal接口的childbirth方法
func (m *mouse) childbirth() {
	fmt.Println("生了一堆小耗子")
}

// 定义鸡结构体
type chicken struct {
}

// 定义speak方法
func (c *chicken) speak() {
	fmt.Println("咯咯咯~~~~~~~")
}
