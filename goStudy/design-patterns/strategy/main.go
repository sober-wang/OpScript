package main

import "fmt"

// Calculate 名为计算的结构题。策略模式中的 Context
type Calculate struct{}

func (cl *Calculate) Plus() {
	fmt.Println("加法")
}

func (cl *Calculate) Subtration() {
	fmt.Println("减法")
}

func (cl *Calculate) CalculusEquation(ce Equation) {
	ce.Answer()

}

// Equation 方程接口，里面有个 Answer() 解答方法；策略模式中的：Stategy
type Equation interface {
	Answer()
}

type Differential struct{}

// Answer() 基于微分方程的具体实现。 策略模式中 ConcreteStrategy
func (*Differential) Answer() {
	fmt.Println("微分方程")
}

type Integral struct{}

func (*Integral) Answer() {
	fmt.Println("积分方程")
}

func main() {
	hard := &Calculate{}
	hard.Plus()
	hard.Subtration()
	hard.CalculusEquation(new(Differential))
	hard.CalculusEquation(new(Integral))

}
