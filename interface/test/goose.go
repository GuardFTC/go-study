// Package test _interface @Author:冯铁城 [17615007230@163.com] 2023-04-03 10:01:21
package test

import "fmt"

// 定义鹅接口
type goose interface {
	releaseSkills()
}

// 定义白板大鹅
type bigGoose struct {
	name string
}

func (b *bigGoose) releaseSkills() {
	fmt.Println(b.name + ":没有技能")
}

// 定义警长
type sheriff struct {
	*bigGoose
}

func (s *sheriff) releaseSkills() {
	fmt.Println(s.name + ":是否刀对了人")
}

// 定义正义使者
type messengerJustice struct {
	*bigGoose
}

func (m *messengerJustice) releaseSkills() {
	fmt.Println(m.name + ":无责刀")
}
