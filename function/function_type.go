// Package function @Author:冯铁城 [17615007230@163.com] 2023-03-14 16:46:05
package function

import "fmt"

// GetFuncType 获取函数类型
func GetFuncType() {
	fmt.Printf("type1's type is %T\n", type1)
	fmt.Printf("type2's type is %T\n", type2)
	fmt.Printf("type3's type is %T\n", type3)
}

func type1() {

}

func type2(a int, b string) {

}

func type3(a int, b string) (c int, d error) {
	return
}
