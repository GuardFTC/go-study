// Package function @Author:冯铁城 [17615007230@163.com] 2023-03-13 19:43:19
package function

import "fmt"

func TestParamMethod() {
	addCAndMinAB(1, 2, 3, getMin)
}

// 基于f对a,b进行操作，并将结果与c相加
func addCAndMinAB(a int, b int, c int, f func(int, int) int) {

	//1.获取a,b最小值
	min := f(a, b)

	//2.与c相加
	sum := min + c

	//3.打印
	fmt.Printf("f's type is %T\nmin(%d,%d) + %d = %d\n", f, a, b, c, sum)
}

// 获取最小值的方法
func getMin(a int, b int) (min int) {
	if a < b {
		min = a
	} else {
		min = b
	}
	return min
}
