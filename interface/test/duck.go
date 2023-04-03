// Package test @Author:冯铁城 [17615007230@163.com] 2023-04-03 10:11:33
package test

import "fmt"

// 定义鸭接口
type duck interface {
	kill()
}

// 定义白板鸭
type bigDuck struct {
	name string
}

func (b *bigDuck) kill() {
	fmt.Println(b.name + ":我只有刀")
}

// 定义刺客
type assassin struct {
	*bigDuck
}

func (a *assassin) kill() {
	fmt.Println(a.name + ":我有刀，我还有枪")
}

// 定义超能力者
type superPowers struct {
	*bigDuck
}

func (s *superPowers) kill() {
	fmt.Println(s.name + ":我感觉自己是个废物")
}
