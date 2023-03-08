// Package datatype @Author:冯铁城 [17615007230@163.com] 2023-03-05 16:15:00
package datatype

import (
	"fmt"
	"time"
)

// PrintTime 时间相关输出操作
func PrintTime() {

	//获取当前时间
	dateTime := time.Now()

	//获取当前UTC时间
	fmt.Printf("now UTC time is %v\n", dateTime.UTC())

	//获取当前时间的年月日、时分秒
	fmt.Printf("year is %d\n", dateTime.Year())
	fmt.Printf("month is %d\n", dateTime.Month())
	fmt.Printf("day is %d\n", dateTime.Day())
	fmt.Printf("hour is %d\n", dateTime.Hour())
	fmt.Printf("minute is %d\n", dateTime.Minute())
	fmt.Printf("second is %d\n", dateTime.Second())

	//获取当前时间秒级时间戳
	fmt.Printf("unix sec timestamp is %d\n", dateTime.Unix())

	//获取当前毫秒时间戳
	fmt.Printf("unix mill timestamp is %d\n", dateTime.UnixMilli())

	//Time类型转为string类型 yyyy-MM-dd HH:mm:ss
	fmt.Printf("datetime to string is %s\n", dateTime.Format(time.DateTime))

	//Time类型转为string类型 yyyy-MM-dd
	fmt.Printf("date to string is %s\n", dateTime.Format(time.DateOnly))

	//Time类型转为string类型 HH:mm:ss
	fmt.Printf("time to string is %s\n", dateTime.Format(time.TimeOnly))

	//string类型转Time
	parseTime, _ := time.Parse(time.DateTime, "2022-02-02 02:02:02")
	fmt.Printf("string to datetime is %v\n", parseTime)
}
