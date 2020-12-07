/*
来源: github.com/i-coder-robot/design-patterns-in-golang
     B站 https://www.bilibili.com/video/BV1GD4y1D7D3/?p=1
*/
package abstract_factory

import "fmt"

/* 出门买东西前，先计划好自己要买的东西，就是定义一个工厂方法 */
// Shopping 购物接口
type Shopping interface {
	Buying()
}

type boots struct {
}

func (b *boots) Buying() {
	fmt.Println("马丁靴")
}

type coat struct {
}

func (c *coat) Buying() {
	fmt.Println("黄色大衣")
}

/*  以下部分列出去哪里买东西，钱多喽，你去商场；钱少，商贸城看看。去哪里买东西重新抽象了买东西的地方  */

// ShoppingInMarket 去购物接口
type ShoppingInMarket interface {
	BuyBoots() Shopping
	BuyCoat() Shopping
}

// GoToMarket 去商场的结构体
type GoToMarket struct {
}

// NewShoppingInMarket 穷的只剩钱了，宝龙，银泰，天街统统买一遍
func NewShoppingInMarket() ShoppingInMarket {
	return &GoToMarket{}
}

// BuyBoots 在商场买鞋子
func (g *GoToMarket) BuyBoots() Shopping {
	return &boots{}
}

// BuyCoat 在商场买大衣
func (g *GoToMarket) BuyCoat() Shopping {
	return &coat{}
}

// NewShoppingInBusiness 生活要节俭；穿不同，暖相同，食不同，饱相同
func NewShoppingInBusiness() ShoppingInMarket {
	return &GoToBusiness{}
}

// GoToBusiness 商场太贵了，我要去商贸城，例如：杭州四季春
type GoToBusiness struct {
}

// BuyBoots 在商贸城买鞋子
func (g *GoToBusiness) BuyBoots() Shopping {
	return &boots{}
}

// BuyCoat 在商贸成买大衣
func (g *GoToBusiness) BuyCoat() Shopping {
	return &coat{}
}
