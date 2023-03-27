// Package map_study @Author:冯铁城 [17615007230@163.com] 2023-03-27 09:44:54
package map_study

import (
	"fmt"
	"sort"
	"strconv"
)

func CreateMap() {

	//1.声明map
	var map1 map[string]int
	var map2 map[string]float64

	//2.初始化map
	map1 = map[string]int{"one": 1, "two": 2}
	map2 = make(map[string]float64)

	//3.设置键值对
	for i := 0; i < 5; i++ {
		map2[strconv.Itoa(i)] = float64(i)
	}

	//4.for循环打印map
	for key, value := range map2 {
		fmt.Printf(" the map2 key is [%s], value is [%f]\n", key, value)
	}

	//5.map为引用类型
	map3 := map1
	map3["three"] = 3

	//6.获取map中的值,并打印
	fmt.Printf("the map1 key=[%s] value=[%v]\n", "one", map1["one"])
	fmt.Printf("the map1 key=[%s] value=[%v]\n", "two", map1["two"])
	fmt.Printf("the map1 key=[%s] value=[%v]\n", "three", map1["three"])
	fmt.Printf("the map3 key=[%s] value=[%v]\n", "three", map3["three"])

	//7.提前声明map的容量
	map4 := make(map[int]int, 5)
	for i := 0; i < 5; i++ {
		map4[i] = i << 2
	}

	//8.校验key值是否存在
	value, isExist := map4[0]
	if isExist {
		fmt.Printf("map4 has key 0, the value is %d\n", value)
	} else {
		fmt.Printf("map4 don't has key 0\n")
	}

	//9.上述的简化写法
	if _, ok := map4[10]; ok {
		fmt.Printf("map4 has key 10, the value is %d\n", value)
	} else {
		fmt.Printf("map4 don't has key 10\n")
	}

	//10.删除key0
	delete(map4, 0)
	if _, ok := map4[0]; ok {
		fmt.Printf("map4 has key 0, the value is %d\n", value)
	} else {
		fmt.Printf("map4 don't has key 0,delete successful\n")
	}

	//11.获取map的长度
	fmt.Printf("the len of map4 is %d\n", len(map4))

	//12.排序map
	//12.1 创建一个无序map
	map5 := map[string]int{"5": 5, "1": 1, "2": 2, "3": 3, "4": 4}
	fmt.Print("not sort map5 is:[ ")
	for key, value := range map5 {
		fmt.Printf("%s:%d ", key, value)
	}
	fmt.Println("]")

	//12.2 创建key切片并排序
	keys := make([]string, 0)
	for key, _ := range map5 {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	fmt.Print("sort map5 is:[ ")
	for _, key := range keys {
		fmt.Printf("%s:%d ", key, map5[key])
	}
	fmt.Println("]")

	//13.键值对对调
	//13.1 创建原map
	map6 := make(map[string]int, 2)
	map6["one"] = 1
	map6["two"] = 2
	for key, value := range map6 {
		fmt.Printf("the map7 key=[%s] value=[%d]\n", key, value)
	}
	//13.2 创建对调map
	map7 := make(map[int]string, len(map6))
	for key, value := range map6 {
		map7[value] = key
	}
	for key, value := range map7 {
		fmt.Printf("the map7 key=[%d] value=[%s]\n", key, value)
	}
}
