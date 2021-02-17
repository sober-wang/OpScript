package builder

import (
	"fmt"
	"testing"
)

func TestConcreteBuilder_GetResult(t *testing.T) {
	// 先创建需要构建的物品
	builder := NewConcreteBuilder()
	// 告诉组装者，我要建什么样的物品
	//director := NewDirector(&builder)
	//director.Construct()
	product := builder.GetResult()
	fmt.Println(product.Built)
}
