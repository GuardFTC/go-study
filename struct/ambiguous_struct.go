// Package _struct @Author:冯铁城 [17615007230@163.com] 2023-03-28 10:42:47
package _struct

import "fmt"

// 声明颜色结构体
type color struct {
	colorName string
	colorDes  string
}

// 声明汽车结构体
type car struct {
	name  string
	price float32
	color
	int
}

func AmbiguousStruct() {

	//1.创建一个汽车结构体
	var car1 car
	car1.name = "别克"
	car1.price = 1000000
	car1.colorName = "黑色"
	car1.colorDes = "高贵的颜色"
	car1.int = 100
	fmt.Printf("the car1 is %v\n", car1)

	//2.通过new创建
	car2 := new(car)
	car2.name = "法拉利"
	car2.price = 5000000
	car2.colorName = "红色"
	car2.colorDes = "青春的红色"
	car2.int = 200
	fmt.Printf("the car2 is %v\n", car2)

	//3.直接创建
	car3 := car{"兰博基尼", 100000000, color{"蓝色", "优雅的蓝色"}, 300}
	fmt.Printf("the car3 is %p\n", &car3)
	fmt.Printf("the car3'color'name is %v\n", car3.color.colorName)
}
