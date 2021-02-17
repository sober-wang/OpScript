package builder

import "fmt"

type Builder interface {
	Build()
}

type Director struct {
	builder Builder
}

func NewDirector(b Builder) Director {
	return Director{builder: b}
}

func (d *Director) Construct() {
	d.builder.Build()
}

type ConcreteBuilder struct {
	built bool
}

func NewConcreteBuilder() ConcreteBuilder {
	return ConcreteBuilder{false}
}

func (c *ConcreteBuilder) Build() {
	c.built = true
}

type Product struct {
	Built bool
}

func (c *ConcreteBuilder) GetResult() Product {
	if c.built {
		fmt.Println("ERROR ")
	}
	return Product{c.built}
}
