// Package _struct @Author:冯铁城 [17615007230@163.com] 2023-03-29 09:45:37
package _struct

import "fmt"

// TestFunction 测试结构体方法
func TestFunction() {

	//1.创建person结构体类型变量
	p := new(person)

	//2.设置名称并自我介绍
	p.setName("马冬梅")
	p.introduce()
	fmt.Printf("the person info is %v\n", p)

	//3.改名
	p.rename("马飞飞")
	fmt.Printf("the person info is %v\n", p)

	//4.偷偷改名介绍
	p.renameAndIntroduce()
	fmt.Printf("the person info is %v\n", p)

	//5.开始走
	p.walk()

	//6.开始跑
	p.run()

	//7.开始说话
	p.speak()

	//8.调用string方法
	fmt.Println("person is", p)
}

// 定义动物结构体
type animal struct{}

// 定义走方法, 接收器为一个animal结构体类型的值接收器
func (animal) walk() {
	fmt.Printf("this one is walking\n")
}

// 定义跑方法
func (animal) run() {
	fmt.Printf("this one is running\n")
}

// 定义猿猴结构体
type ape struct{}

// 定义说话方法
func (a *ape) speak() {
	fmt.Printf("this one is speaking\n")
}

// 定义人类结构体
type person struct {
	name string
	animal
	ape
}

// setName
func (p *person) setName(name string) {
	p.name = name
}

// getName
func (p *person) getName() (name string) {
	return p.name
}

// 重新命名
func (p *person) rename(name string) {

	//1.打印原名称
	fmt.Printf("At last my name is %s\n", p.getName())

	//2.改名
	p.setName(name)

	//3.打印现名称
	fmt.Printf("Now my name is %s\n", p.getName())
}

// 自我介绍 接收器建议只使用指针接收器或值接收器当中的一种,不建议混合使用
func (p person) introduce() {
	fmt.Printf("My name is %s\n", p.getName())
}

// 偷偷改名并自我介绍
func (p person) renameAndIntroduce() {

	//偷偷改个名称
	p.setName("王大锤")
	fmt.Printf("My name is %s\n", p.getName())
}

func (p person) String() string {
	return "[" + p.name + "]"
}

func CreatePerson(name string) *person {
	return &person{name: name}
}
