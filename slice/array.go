// Package slice @Author:冯铁城 [17615007230@163.com] 2023-03-15 14:58:48
package slice

import (
	"fmt"
)

// CreatArray 创建普通数组
func CreatArray() {

	//1.创建一个数组
	var array [3]int

	//2.遍历数组
	for i := 0; i < len(array); i++ {
		fmt.Printf("value=%d ", array[i])
	}
	fmt.Println()

	//3.遍历数组
	for index, value := range array {
		fmt.Printf("index=%d value=%d\n", index, value)
	}

	//4.数组常量
	var constArray = [3]int{1, 2, 3}
	for index, value := range constArray {
		fmt.Printf("index=%d value=%d\n", index, value)
	}

	//5.数组为值类型
	array = [3]int{1, 2, 3}
	var copyArray = array
	copyArray[2] = 4
	fmt.Printf("array memory is %p\n", &array)
	fmt.Printf("copyArray memory is %p\n", &copyArray)
	fmt.Printf("array[2]=%d copyArray[2]=%d\n", array[2], copyArray[2])

	//7.多维数组
	array2D := [2][2]int{{1, 2}, {3, 4}}
	fmt.Printf("array2D=%v\n", array2D)

	//8.数组指针作为入参
	array = [3]int{11, 12, 13}
	ptr := &array
	arrayItemDouble(ptr)
	for index, value := range array {
		fmt.Printf("index=%d value=%d\n", index, value)
	}
}

// MemoryArray 数组的内存分析
func MemoryArray() {

	//1.创建一个数组
	array := [3]int{1, 2, 3}

	//2.输出数组引用的内存地址
	fmt.Printf("array location is %p\n", &array)

	//3.输出数组内部各个元素的内存地址
	for i := 0; i < len(array); i++ {
		fmt.Printf("array[%d] location is %p\n", i, &array[i])
	}
}

// 数组元素*2
func arrayItemDouble(ptr *[3]int) {
	for i := 0; i < len(*ptr); i++ {
		(*ptr)[i] = (*ptr)[i] * 2
	}
}
