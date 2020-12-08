/*
来源: github.com/i-coder-robot/design-patterns-in-golang
     B站 https://www.bilibili.com/video/BV1GD4y1D7D3/?p=1

设计模式学习
3. 门面模式 

学习总结：
我们平时使用的 RESTful API ,web 服务页面接口等就可以用门面模式实现，每一条 url 文根就是门面中的某一类物品，例如本例中买糖糕

*/
package facade

import "fmt"

// Sweet 甜品窗口
type Sweet struct{
}

// NewSweet 初始化甜品窗口，看看有什么甜品
func NewSweet() *Sweet{
	return &Sweet{}
}

// SweetCake 我喜欢吃新丰小吃的糖糕
func (xf *Sweet) SweetCake() {
	fmt.Println("老板请给我一个糖糕；老板开始制作糖糕... ；获得糖糕，吃起来")
}

// MaYuan 麻元列表
type MaYuan struct{
}

// NewMaYuan 初始化麻元
func NewMaYuan() *MaYuan{
	return &MaYuan{}
}

// GetMaYuan 麻元虽然油腻，但是抵饿
func (my *MaYuan) GetMaYuan() {
	fmt.Println("老板请给我一个麻元；老板取出一个麻元...；获得麻元，吃起来")
}

// Meet 一餐总的有肉吧
type Meet struct{
}

// NewMeet 初始化肉类窗口
func NewMeet() *Meet{
	return &Meet{}
}

// GetChickenLeg 加个鸡腿把
func (cl *Meet) GetChickenLeg(){
	fmt.Println("老板请给我一个鸡腿；老板取出一个鸡腿...；获得鸡腿，吃起来")
}

// XinFengWindow 新丰小吃的窗口
type XinFengWindow struct {
	buySweet *Sweet
	buyMayuan *MaYuan
	buyMeet *Meet
}

// NewXinFengWindow 到达新丰小吃的窗口
func NewXinFengWindow() *XinFengWindow{
	return &XinFengWindow{
		buySweet: NewSweet(),
		buyMayuan: NewMaYuan(),
		buyMeet: NewMeet(),
	}
}

// Buy 买小吃啦！
func (nxfw *XinFengWindow) Buy(){
	fmt.Println("挑选商品....")
	nxfw.buySweet.SweetCake()
	nxfw.buyMayuan.GetMaYuan()
	nxfw.buyMeet.GetChickenLeg()
	fmt.Println("离开店铺....")
}

