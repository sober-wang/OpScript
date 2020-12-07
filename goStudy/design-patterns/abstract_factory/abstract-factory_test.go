package abstract_factory

import (
	"fmt"
	"testing"
)

func TestGoToMarket(t *testing.T) {
	market := NewShoppingInMarket()

	boots := market.BuyBoots()
	boots.Buying()

	coat := market.BuyCoat()
	coat.Buying()

	fmt.Println("-------------- 楚河汉界 ------------------")
	business := NewShoppingInBusiness()
	bb := business.BuyBoots()
	bb.Buying()

	bc := business.BuyCoat()
	bc.Buying()

}
