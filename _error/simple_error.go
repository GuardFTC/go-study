// Package _error @Author:冯铁城 [17615007230@163.com] 2023-04-13 09:59:32
package _error

import (
	"errors"
	"fmt"
	"strconv"
)

// CustomError 自定义异常
type CustomError struct {
	msg  string
	code int
}

// 实现Error方法
func (e CustomError) Error() string {
	return fmt.Sprintf("custom error msg is [%s],code is [%d]", e.msg, e.code)
}

// SimpleError 异常的基础操作
func SimpleError() {

	//1.使用errors包声明一个异常
	err := errors.New("errors test error")
	fmt.Printf("error is %v,type is %T\n", err.Error(), err)

	//2.使用fmt创建异常
	err = fmt.Errorf("error is %v", "fmt test error")
	fmt.Printf("%v,type is %T\n", err.Error(), err)

	//3.自定义异常
	customError := new(CustomError)
	customError.msg = "custom error!!!"
	customError.code = 200
	fmt.Println(customError.Error())
}

// PanicAndRecoverAndDefer panic、defer、recover学习
func PanicAndRecoverAndDefer(id int) {
	defer func() {
		fmt.Println("PanicAndRecoverAndDefer is end")
		if err := recover(); err != nil {
			fmt.Println("some error was happened")
			fmt.Println(err)
		}
	}()

	if id <= 0 {
		panic("the id is not a right value")
	} else {
		fmt.Printf("the id's value is %v\n", id)
	}
}

// PanicAndError 包外部方法
func PanicAndError(id int) (res string, err error) {

	//1.定义defer逻辑
	defer func() {
		if e := recover(); e != nil {
			error, ok := e.(error)
			if ok {
				err = CustomError{error.Error(), 400}
			} else {
				err = CustomError{fmt.Errorf("%v", e).Error(), 500}
			}
		}
	}()

	//2.调用内部方法
	res = panicAndError(id)

	//3.返回
	return
}

// TryUnitTest Golang中的单元测试示例
func TryUnitTest(num int) (isLessZero bool) {
	if num < 0 {
		isLessZero = true
	} else {
		isLessZero = false
	}
	return
}

// 包内部方法
func panicAndError(id int) (name string) {
	if id > 0 {
		return "ftc" + strconv.Itoa(id)
	} else {
		panic("the id is not a right value")
	}
}
