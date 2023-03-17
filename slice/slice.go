// Package slice @Author:冯铁城 [17615007230@163.com] 2023-03-16 10:40:41
package slice

import (
	"fmt"
	"strconv"
)

// CreateSlice 创建切片
func CreateSlice() {

	//1.根据数组生成切片
	array := [5]int{0, 1, 2, 3, 4}
	slice1 := array[0:len(array)]
	printSlice(slice1)
	slice2 := array[0:]
	printSlice(slice2)
	slice3 := array[:len(array)]
	printSlice(slice3)
	slice4 := array[:]
	printSlice(slice4)

	//2.根据数组生成不同长度的切片
	slice5 := array[0:3]
	printSlice(slice5)
	slice6 := array[1:3]
	printSlice(slice6)

	//3.直接生成切片
	slice7 := []int{0, 1, 2, 3, 4}
	printSlice(slice7)

	//4.基于切片生成切片
	slice7 = slice7[2:4]
	printSlice(slice7)
	slice7 = slice7[1:2]
	printSlice(slice7)

	//5.用make一个切片
	slice8 := make([]int, 5, 10)
	printSlice(slice8)

	//6.用new创建一个切片
	slice9 := new([10]int)[0:5]
	printSlice(slice9)

	//7.遍历切片fori
	slice10 := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(slice10); i++ {
		fmt.Printf("index=%d value=%d\n", i, slice10[i])
	}

	//8.遍历切片forr
	for index, value := range slice10 {
		fmt.Printf("index=%d value=%d\n", index, value)
	}

	//9.复制切片,new < old
	oldSlice := []int{1, 2, 3, 4}
	newSlice := make([]int, 1, 2)
	copyItemNum := copy(newSlice, oldSlice)
	fmt.Println("copy success item num = " + strconv.Itoa(copyItemNum))
	printSlice(newSlice)

	//10.复制切片,new > old
	oldSlice = []int{1, 2, 3, 4}
	newSlice = make([]int, 5, 10)
	copyItemNum = copy(newSlice, oldSlice)
	fmt.Println("copy success item num = " + strconv.Itoa(copyItemNum))
	printSlice(newSlice)

	//11.向切片添加元素
	newSlice = make([]int, 1)
	printSlice(newSlice)
	newSlice = append(newSlice, 24, 30)
	printSlice(newSlice)

	//12.向切片添加切片
	newSlice = append(newSlice, oldSlice...)
	printSlice(newSlice)

	//13.复制指定切片
	copySlice := make([]int, len(newSlice))
	copy(copySlice, newSlice)
	printSlice(copySlice)

	//14.切片是引用类型
	slice := []int{1, 2, 3}
	copySlice = slice
	copySlice[1] = 100
	fmt.Printf("slice memory is %p\n", &slice)
	fmt.Printf("copySlice memory is %p\n", &copySlice)
	fmt.Printf("slice[1]=%d copySlice[1]=%d\n", slice[1], copySlice[1])
}

// UseAppend 使用append函数
func UseAppend() {

	//1.创建切片
	slice := make([]int, 0)
	printSlice(slice)

	//2.向切片中添加元素
	slice = append(slice, 1, 2, 3)
	printSlice(slice)

	//3.移除第2个元素
	slice = append(slice[:1], slice[2:]...)
	printSlice(slice)

	//4.在第二个位置append一个新的切片/元素
	newSlice := []int{30, 23, 35, 11}
	slice = append(slice[:1], append(newSlice, slice[1:]...)...)
	printSlice(slice)

	//5.取出位于切片的最末尾元素
	endItem := slice[len(slice)-1:][0]
	fmt.Printf("the end item is %d", endItem)
}

// MemorySlice Slice内存分析
func MemorySlice() {

	//1.创建一个数组
	array := [3]int{1, 2, 3}

	//2.输出数组引用的内存地址
	fmt.Printf("array location is %p\n", &array)

	//3.输出数组内部各个元素的内存地址
	for i := 0; i < len(array); i++ {
		fmt.Printf("array[%d] location is %p\n", i, &array[i])
	}
	fmt.Println("--------------------------------------------------------------------")

	//4.创建一个slice
	slice := array[:]

	//5.输出切片引用的内存地址
	fmt.Printf("slice location is %p\n", &slice)

	//6.输出切片指向堆内存的位置
	fmt.Printf("slice heap location is %p\n", slice)

	//7.输出数组内部各个元素的内存地址
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%d] location is %p\n", i, &slice[i])
	}
	fmt.Println("--------------------------------------------------------------------")

	//8.slice扩容
	slice = append(slice, 4, 5)

	//9.输出切片引用的内存地址
	fmt.Printf("slice location is %p\n", &slice)

	//10.输出切片指向堆内存的位置
	fmt.Printf("slice heap location is %p\n", slice)

	//11.输出数组内部各个元素的内存地址
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%d] location is %p\n", i, &slice[i])
	}
}

// EnlargeSlice Slice扩容机制
func EnlargeSlice() {

	//1.规则验证
	slice := make([]int, 0)
	oldCap := cap(slice)
	for i := 0; i < 2048; i++ {
		slice = append(slice, i)
		newCap := cap(slice)
		if oldCap != newCap {
			fmt.Printf(" slice memory is %p, pointer memory is %p, cap is %d\n", &slice, slice, cap(slice))
			oldCap = newCap
		}
	}
	fmt.Println("--------------------------------------------------------------")

	//2.新老数组验证
	slice = make([]int, 0, 3)
	fmt.Printf(" slice memory is %p, pointer memory is %p, len is %d, cap is %d\n", &slice, slice, len(slice), cap(slice))
	for i := 0; i < 4; i++ {
		slice = append(slice, 1)
		fmt.Printf(" slice memory is %p, pointer memory is %p, len is %d, cap is %d\n", &slice, slice, len(slice), cap(slice))
	}
}

// 输出切片详情
func printSlice(slice []int) {
	fmt.Printf("slice=%v len=%d cap=%d\n", slice, len(slice), cap(slice))
}
