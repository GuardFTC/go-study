// Package datatype @Author:冯铁城 [17615007230@163.com] 2023-03-05 09:34:15
package datatype

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// 定义字符串
var str1 = "ftc"

// PrintString 输出字符串
func PrintString() {

	//输出字符串的长度
	fmt.Printf("%s length is %d\n", str1, len(str1))

	//比较字符串
	str2 := "ftc"
	fmt.Printf("str1 == str2?:%t\n", str1 == str2)

	//输出字符串的字符
	fmt.Printf("the first char at %s is %c\n", str2, str2[0])
	fmt.Printf("the second char at %s is %c\n", str2, str2[1])
	fmt.Printf("the last char at %s is %c\n", str2, str2[len(str2)-1])

	//截取字符串
	fmt.Printf("%s sub 1-2 is %s\n", str2, str2[0:2])

	//截取非ASCII编码字符串
	strC := "我喜欢库里，他妈的爱谁谁"
	runes := []rune(strC)
	fmt.Printf("%s sub 1-4 is %s\n", strC, string(runes[0:4]))

	//输出非ASCII编码字符串长度
	fmt.Printf("%s length is %d\n", strC, len(strC))
	fmt.Printf("%s length is %d\n", strC, utf8.RuneCountInString(strC))

	//判定前缀后缀
	str3 := "ftc is a cool boy"
	fmt.Printf("%s start with ftc?:%t\n", str3, strings.HasPrefix(str3, str2))
	fmt.Printf("%s end with ftc?:%v\n", str3, strings.HasSuffix(str3, str2))

	//字符串包含
	str4 := "cool boy"
	fmt.Printf("%s contains %s?:%t\n", str3, str4, strings.Contains(str3, str4))

	//字符串出现的位置
	fmt.Printf("%s show in %s the first index is %d\n", "c", str3, strings.Index(str3, "c"))
	fmt.Printf("%s show in %s the first index is %d\n", "is", str3, strings.Index(str3, "is"))
	fmt.Printf("%s show in %s the last index is %d\n", "c", str3, strings.LastIndex(str3, "c"))

	//字符串替换
	str5 := "curry is so good! I love curry so much!"
	fmt.Println(strings.Replace(str5, "curry", "klay", -1))

	//统计字符串出现次数
	fmt.Printf("curry show in %s %d times\n", str5, strings.Count(str5, "curry"))

	//重复字符串
	str6 := "重要的事情说三遍！"
	fmt.Println(strings.Repeat(str6, 3))

	//大小写转换
	str7 := strings.ToLower("JAMES HARDEN")
	fmt.Printf("%s to lower case is %s\n", "JAMES HARDEN", str7)
	fmt.Printf("%s to upper case is %s\n", str7, strings.ToUpper(str7))

	//按空格切分字符串
	str8 := "who can stop me!!!!!"
	fields := strings.Fields(str8)
	for _, item := range fields {
		fmt.Println(item)
	}

	//自定义字符切分字符串
	str9 := "i-like-curry-and-klay"
	fields = strings.Split(str9, "-")
	for _, item := range fields {
		fmt.Println(item)
	}

	//将字符数组拼接为字符串
	fmt.Println(strings.Join(fields, "|"))

	//int类型与字符串类型互换
	fmt.Printf("66666 to string is %s\n", strconv.Itoa(66666))
	atoi, _ := strconv.Atoi("66666")
	fmt.Printf("66666 to int is %d\n", atoi)

	//其他类型转字符串
	fmt.Printf("true to string is %s\n", strconv.FormatBool(true))
	fmt.Printf("66666 to string is %s\n", strconv.FormatInt(66666, 10))
	fmt.Printf("88888 to string is %s\n", strconv.FormatUint(88888, 10))
	fmt.Printf("3.1415926 to string is %s\n", strconv.FormatFloat(3.1415926, 'f', 7, 64))

	//字符串转其他类型
	parseBool, _ := strconv.ParseBool("true")
	fmt.Printf("true to boolean is %t\n", parseBool)
	parseInt, _ := strconv.ParseInt("66666", 10, 32)
	fmt.Printf("66666 to int is %d\n", parseInt)
	parseUint, _ := strconv.ParseUint("88888", 10, 32)
	fmt.Printf("88888 to uint is %d\n", parseUint)
	parseFloat, _ := strconv.ParseFloat("3.1415926", 64)
	fmt.Printf("3.1415926 to float is %f\n", parseFloat)
}
