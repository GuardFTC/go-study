// Package _interface @Author:冯铁城 [17615007230@163.com] 2023-04-02 17:51:41
package _interface

import (
	"fmt"
	"reflect"
)

// Ref 反射相关方法
func Ref() {

	//1.创建int类型变量
	i := 1

	//2.通过反射获取类型
	t := reflect.TypeOf(i)
	fmt.Printf("the i's type is %v\n", t.Name())

	//3.通过反射获取值
	v := reflect.ValueOf(i)
	fmt.Printf("the i's value is %v\n", v.Int())

	//4.通过v可以获取到类型
	fmt.Printf("the v's kind is %v\n", v.Kind().String())

	//5.通过反射设置值-失败示例
	if ok := v.CanSet(); ok {
		v.SetInt(10)
		fmt.Printf("the v's value is %v\n", v.Int())
	} else {
		fmt.Printf("can't update v's value\n")
	}

	//6.通过反射设置值-成功示例
	v = reflect.ValueOf(&i).Elem()
	if ok := v.CanSet(); ok {
		v.SetInt(10)
		fmt.Printf("the v's value is %v\n", v.Int())
	} else {
		fmt.Printf("can't update v's value\n")
	}

	//7.Elem方法的作用-获取指针指向内存地址的值
	fmt.Printf("reflect.ValueOf(&refStruct)'type = %v\n", reflect.ValueOf(&i).Kind())
	fmt.Printf("reflect.ValueOf(&refStruct).Elem()'type = %v\n", reflect.ValueOf(&i).Elem().Kind())
}

// RefStruct 反射使用结构体
type RefStruct struct {
	s1, s2, s3 string
	S          string
}

// ConsoleFields ref结构体方法,值变量接收器
func (r RefStruct) ConsoleFields() {
	fmt.Printf("s1=[%v] s2=[%v] s3=[%v]\n", r.s1, r.s2, r.s3)
}

// ConsoleFieldsByPointer ref结构体方法,指针变量接收器
func (r *RefStruct) ConsoleFieldsByPointer() {
	fmt.Printf("s1‘pointer=[%v] s2‘pointer=[%v] s3‘pointer=[%v]\n", r.s1, r.s2, r.s3)
}

// StructRef 结构体反射相关方法
func StructRef() {

	//1.创建结构体类型值变量
	refStruct := RefStruct{"ftc", "zyl", "gn", "skx"}

	//2.通过反射获取结构体变量字段类型相关，以及字段值
	t := reflect.TypeOf(refStruct)
	v := reflect.ValueOf(refStruct)
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)
		fmt.Printf("the field name is %v, type is %v, value is %v\n", field.Name, field.Type, fieldValue.String())
	}

	//3.通过反射修改结构体变量字段，只能修改导出字段（首字母大写字段）
	v = reflect.ValueOf(&refStruct).Elem()
	fmt.Printf("before update, the field value is %v\n", v.FieldByName("S").String())
	if ok := v.CanSet(); ok {
		v.FieldByName("S").SetString("Upper")
	}
	fmt.Printf("after update, the field value is %v\n", v.FieldByName("S").String())

	//4.通过反射调用方法,值类型接收器,只能反射导出方法（首字母大写）
	m := reflect.ValueOf(refStruct).MethodByName("ConsoleFields")
	if m.IsValid() {
		m.Call(nil)
	} else {
		fmt.Printf("no method ConsoleFields at RefStruct\n")
	}

	//5.通过反射调用方法,指针类型接收器,只能反射导出方法（首字母大写）
	m = reflect.ValueOf(&refStruct).MethodByName("ConsoleFieldsByPointer")
	if m.IsValid() {
		m.Call(nil)
	} else {
		fmt.Printf("no method ConsoleFieldsByPointer at RefStruct\n")
	}
}
