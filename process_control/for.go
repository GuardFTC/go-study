// Package process_control @Author:冯铁城 [17615007230@163.com] 2023-03-08 13:42:47
package process_control

import "fmt"

func For() {

	//1.最简单的for循环
	for i := 0; i < 5; i++ {
		fmt.Printf("%d,", i)
	}
	fmt.Println()

	//2.条件判定for
	a := 5
	for a > 0 {
		fmt.Printf("%d,", a)
		a--
	}
	fmt.Println()

	//3.range-for
	for key, value := range "abcde" {
		fmt.Printf("key=%c value=%c\n", key, value)
	}

	//4.break
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Printf("%d,", i)
	}
	fmt.Println()

	//5.continue
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Printf("%d,", i)
	}
	fmt.Println()

	//6.label与goto
	a = 1
HERE:
	fmt.Printf("%d,", a)
	if a == 3 {
		return
	}
	a++
	goto HERE
}
