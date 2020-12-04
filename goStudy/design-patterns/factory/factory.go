/*
来源: github.com/i-coder-robot/design-patterns-in-golang
     B站 https://www.bilibili.com/video/BV1GD4y1D7D3/?p=1

设计模式学习
1 工厂模式

*/

package factory

import "fmt"

// Canteen 食堂接口
type Canteen interface {
	MakeFood(x string)
}

type Soup struct {
}

func (soup *Soup) MakeFood(x string) {
	switch x {
	case "x":
		fmt.Println("西红柿蛋汤")
	case "g":
		fmt.Println("骨头汤")
	default:
		fmt.Println("刷锅水")
	}
	//fmt.Println("刷锅水")
}

type MainCourse struct {
}

func (mc *MainCourse) MakeFood(x string) {
	switch x {
	case "meats":
		fmt.Println("红烧排骨")
	case "vegetables":
		fmt.Println("醋溜白菜")
	default:
		fmt.Println("黄土一把")
	}
	//fmt.Println("黄土一把")
}

func NewCateen(s string) Canteen {
	switch s {
	case "tang":
		return &Soup{}
	case "cai":
		return &MainCourse{}
	}
	return nil

}
