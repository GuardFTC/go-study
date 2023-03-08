// Package process_control @Author:冯铁城 [17615007230@163.com] 2023-03-08 10:58:00
package process_control

import "fmt"

// SwitchCase switch-case流程的各种写法
func SwitchCase() {

	//1.最简单写法
	num := 1
	switch num {
	case 1:
		fmt.Println("run the case 1 process")
	case 2:
		fmt.Println("run the case 2 process")
	case 3:
		fmt.Println("run the case 3 process")
	default:
		fmt.Println("run the default process")
	}

	//2.带有fallthrough的switch-case
	num = 2
	switch num {
	case 1:
		fmt.Println("run the case 1 process")
	case 2:
		fmt.Println("run the case 2 process")
		fallthrough
	case 3:
		fmt.Println("run the case 3 process")
	default:
		fmt.Println("run the default process")
	}

	//3.不指定switch选项写法
	num = 3
	switch {
	case num == 1:
		fmt.Println("run the case 1 process")
	case num == 2:
		fmt.Println("run the case 2 process")
	case num == 3:
		fmt.Println("run the case 3 process")
	default:
		fmt.Println("run the default process")
	}

	//4.通过方法返回值判定
	switch test1(24) {
	case 1:
		fmt.Println("run the case 1 process")
	case -1:
		fmt.Println("run the case -1 process")
	default:
		fmt.Println("run the default process")
	}

	//5.通过方法返回值的二次计算结果判定
	switch i, j := test2(2); {
	case i == j:
		fmt.Println("run the case i==j process")
	case i != j:
		fmt.Println("run the case i!=j process")
	default:
		fmt.Println("run the default process")
	}
}

// 测试方法
func test1(i int) int {
	if i > 0 {
		return 1
	} else {
		return -1
	}
}

// 测试方法
func test2(i int) (int, int) {
	if i > 0 {
		return 1, 1
	} else {
		return -1, 1
	}
}
