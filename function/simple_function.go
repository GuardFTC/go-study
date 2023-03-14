// Package function @Author:冯铁城 [17615007230@163.com] 2023-03-13 11:37:49
package function

import (
	"fmt"
	"time"
)

func TestMax() {

	//1.调用非命名返回值
	fmt.Printf("getMax param:[a=%d,b=%d] return value:[%d]\n", 1, 2, getMax(1, 2))

	//2.调用命名返回值
	fmt.Printf("getMax2 param:[a=%d,b=%d] return value:[%d]\n", 3, 2, getMax2(3, 2))

	//3.调用多返回值
	max, _ := getMax3(4, 5)
	fmt.Printf("getMax3 param:[a=%d,b=%d] return value:[%d]\n", 4, 5, max)

	//4.调用直接改变参数方法
	ptr := &max
	getMax4(8, 7, ptr)
	fmt.Printf("getMax4 param:[a=%d,b=%d] return value:[%d]\n", 8, 7, max)

	//5.调用变长参数参数
	fmt.Printf("getMax5 param:[a=%d,b=%d] return value:[%d]\n", 11, 12, getMax5(11, 12))
	nums := []int{3, 4, 5, 6, 7, 8}
	max = getMax5(nums...)
	fmt.Printf("getMax5 param:[%v] return value:[%d]\n", nums, max)

	//6.获取方法执行时间
	consoleFuncExecutionTime()
}

// 获取最大值,非命名返回值
func getMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 获取最大值，命名返回值
func getMax2(a int, b int) (max int) {
	if a > b {
		max = a
	} else {
		max = b
	}
	return max
}

// 获取最大值,多返回值
func getMax3(a int, b int) (max int, err error) {
	if a > b {
		max = a
		err = nil
	} else {
		max = b
		err = nil
	}
	return max, err
}

// 获取最大值，直接改变参数
func getMax4(a int, b int, max *int) {
	if a > b {
		*max = a
	} else {
		*max = b
	}
}

// 获取最大值，参数为变长参数
func getMax5(nums ...int) (max int) {
	max = nums[0]
	for _, value := range nums {
		if value > max {
			max = value
		}
	}
	return max
}

// 控制台打印方法执行时间
func consoleFuncExecutionTime() {

	//1.获取方法开始时间
	start := time.Now()

	//2.执行业务逻辑
	for i := 0; i < 1000000; i++ {
		fmt.Print()
	}

	//3.获取方法执行结束时间
	end := time.Now()

	//4.打印总执行时间
	fmt.Printf("this func cost time is %v", end.Sub(start))
}
