/*
来源: github.com/i-coder-robot/design-patterns-in-golang
     B站 https://www.bilibili.com/video/BV1GD4y1D7D3/?p=5
*/

package Bridage

import (
	"fmt"
	"time"
)

type Draw interface {
	DrawRectangle(x, y int)
}

type RedRectangle struct {
}

func (rr *RedRectangle) DrawRectangle(x, y int) {
	fmt.Printf("The Red recatangle long: %v; The Red recatangle wide: %v \n", x, y)
}

type BlueRectangle struct {
}

func (br *BlueRectangle) DrawRectangle(x, y int) {
	fmt.Printf("The Blue recatangle long: %v; The Blue recatangle wide: %v \n", x, y)
}

// Build 是下方 Rectangle 与上方画法 draw 之间的桥梁，链接了需求和具体实现。
type Build struct {
	draw Draw
}

func (b *Build) Build(d Draw) {
	b.draw = d
	time.Now().Unix()
}

type Rectangle struct {
	build Build
	x     int
	y     int
}

func (r *Rectangle) Constructor(x, y int, draw Draw) {
	r.x = x
	r.y = y
	r.build.Build(draw)
}

func (r *Rectangle) Do() {
	r.build.draw.DrawRectangle(r.x, r.y)
}
