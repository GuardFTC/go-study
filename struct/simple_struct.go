// Package _struct @Author:冯铁城 [17615007230@163.com] 2023-03-28 09:27:57
package _struct

import (
	"fmt"
	"reflect"
)

// 声明一个结构体类型
type student struct {
	id   int    "the student's id"
	name string "the student's name"
	age  int    "the student's age"
}

// 结构体别名
type youngMan student

// CreateStruct 创建结构体
func CreateStruct() {

	//1.直接创建student变量
	var student1 student
	fmt.Printf("the student1 is %v\n", student1)
	student1.id = 1
	student1.name = "ftc"
	student1.age = 27
	fmt.Printf("the student1 is %v\n", student1)

	//2.通过new初始化
	var student2 = new(student)
	fmt.Printf("the student2 is %v\n", student2)
	student2.id = 2
	student2.name = "zyl"
	student2.age = 24
	fmt.Printf("the student2 is %v\n", student2)

	//3.结构体变量直接赋值
	var student3 = student{3, "gn", 26}
	fmt.Printf("the student3 is %v\n", student3)
	student3.age = 29
	fmt.Printf("the student3 is %v\n", student3)

	//4.结构体指针直接赋值
	var student4 = &student{4, "skx", 24}
	fmt.Printf("the student4 is %v\n", student4)
	student4.age = 18
	fmt.Printf("the student4 is %v\n", student4)

	//5.通过结构体别名创建对象
	var youngMan1 youngMan
	fmt.Printf("the youngMan1 is %v\n", youngMan1)
	youngMan1.id = 11
	youngMan1.name = "wqw"
	youngMan1.age = 18
	fmt.Printf("the youngMan1 is %v\n", youngMan1)

	//6.工厂方法模式创建一个结构体
	student5 := newStudent(5, "wkf", 18)
	fmt.Printf("the student5 is %v\n", student5)
	student6 := newStudentBak(6, "nn", 26)
	fmt.Printf("the student6 is %v\n", student6)
}

// StructMemory 结构体内存分析
func StructMemory() {

	//1.声明一个结构体类型
	var student1 student
	fmt.Printf("the memory of student1 is %p\n", &student1)
	fmt.Printf("the memory of student1 is %p\n", &student1.id)
	fmt.Printf("the memory of student1 is %p\n", &student1.name)
	fmt.Printf("the memory of student1 is %p\n", &student1.age)
	fmt.Println("-------------------------------------------------")

	//2.对student1赋值
	student1.id = 1
	student1.name = "ftc"
	student1.age = 27
	fmt.Printf("the memory of student1 is %p\n", &student1)
	fmt.Printf("the memory of student1 is %p\n", &student1.id)
	fmt.Printf("the memory of student1 is %p\n", &student1.name)
	fmt.Printf("the memory of student1 is %p\n", &student1.age)
	fmt.Println("-------------------------------------------------")

	//3.通过new创建一个结构体
	student2 := new(student)
	fmt.Printf("the memory of student2 is %p\n", &student2)
	fmt.Printf("the memory of student2'ptr is %p\n", student2)
	fmt.Printf("the memory of student1 is %p\n", &student2.id)
	fmt.Printf("the memory of student1 is %p\n", &student2.name)
	fmt.Printf("the memory of student1 is %p\n", &student2.age)
	fmt.Println("-------------------------------------------------")

	//4.对student2赋值
	student1.id = 2
	student1.name = "zyl"
	student1.age = 26
	fmt.Printf("the memory of student2 is %p\n", &student2)
	fmt.Printf("the memory of student2'ptr is %p\n", student2)
	fmt.Printf("the memory of student1 is %p\n", &student2.id)
	fmt.Printf("the memory of student1 is %p\n", &student2.name)
	fmt.Printf("the memory of student1 is %p\n", &student2.age)
	fmt.Println("-------------------------------------------------")
}

// StructTag 结构体标签操作
func StructTag() {

	//1.创建一个student对象
	student1 := new(student)

	//2.ref操作
	typeStudent := reflect.TypeOf(*student1)

	//3.遍历获取student类型的字段
	for i := 0; i < typeStudent.NumField(); i++ {
		field := typeStudent.Field(i)
		fmt.Printf("the field %v'tag is [%v]\n", field.Name, field.Tag)
	}
}

// 工厂方法创建一个student
func newStudent(id int, name string, age int) *student {
	return &student{id, name, age}
}

// 工厂方法创建一个student
func newStudentBak(id int, name string, age int) student {
	return student{id, name, age}
}
