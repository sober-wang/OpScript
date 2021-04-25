package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack struct {
	MaxTop int
	Top    int
	List   [10]int
}

func CreateStack(stackDeep int) *Stack {
	return &Stack{
		MaxTop: stackDeep,
		Top:    -1,
	}
}

func (s *Stack) IsNil() bool {
	if s.Top == -1 {
		return true
	}
	return false
}
func (s *Stack) IsFull() bool {
	if s.Top == s.MaxTop-1 {
		return true
	}
	return false
}

func (s *Stack) Push(value int) error {
	if s.IsFull() {
		fmt.Printf("栈满了，元素[ %v ] 添加失败\n", value)
		return errors.New("栈满了")
	}

	s.Top++
	s.List[s.Top] = value
	return nil
}

func (s *Stack) Pop() (int, error) {
	var retVal int
	if s.IsNil() {
		fmt.Println("栈是空的")
		return retVal, errors.New("栈是空的")
	}

	fmt.Printf("元素 [ %v ] 出栈\n", s.List[s.Top])
	//s.List[s.Top] = 0
	retVal = s.List[s.Top]
	s.Top--
	return retVal, nil
}

func (s *Stack) Show() {
	if s.IsNil() {
		fmt.Println("栈是空的")
		return
	}

	for i := s.Top; i > -1; i-- {
		fmt.Printf("栈中第 [ %v ] = [ %v ]\n", i, s.List[i])
	}
}

// ExprCount 算数计算，传入两个需要计算的数，exprOper 是算数计算符
func (s *Stack) ExprCount(one, tow int, exprOper int) int {
	var resul int
	switch {
	// 这里要注意，先出栈的数是在运算符号的右侧，例如：除法中先出栈的就是除数，后出栈的就是被除数。
	case exprOper == 43:
		resul = tow + one
	case exprOper == 45:
		resul = tow - one
	case exprOper == 42:
		resul = tow * one
	case exprOper == 47:
		resul = tow / one
	default:
		fmt.Println("运算符错误")
	}
	return resul
}

func (s *Stack) IsOper(val int) bool {
	if val == 43 || val == 45 || val == 42 || val == 47 {
		return true
	}
	return false
}

// IsExprOperLevel 判断运算符，同时定级
// + - 为 1
// * / 为 2
func (s *Stack) IsExprOperLevel(val int) int {
	var retVal int
	// 43 == + ，45 == - 号
	if val == 43 || val == 45 {
		retVal = 1
	} else if val == 42 || val == 47 {
		retVal = 2
	}
	return retVal
}

func exprNumber(s string) {
	fmt.Printf("计算一下这个式子： %v\n", s)
	numberStack := CreateStack(10)
	exprStack := CreateStack(10)

	var (
		// 滑动值用于扫描算数表达式使用
		index int
		// 数栈中弹出的第一数
		one int
		// 数栈中弹出的第二个数
		tow int
		// 运算后结果保存
		tmp int
		// 运算符 ascii 码存储
		opr int
		// 定义一个合并多位数的变量
		joinNumber string
	)
	for {
		// 取出第一个值
		c := s[index : index+1]
		// 获取当前字符的 ascii 码
		ascNum := int([]byte(c)[0])
		// 判断当前字符是否是运算符
		fmt.Printf("String: %v; ASCII: %v\n", c, ascNum)
		if exprStack.IsOper(ascNum) {
			if exprStack.Top == -1 {
				exprStack.Push(ascNum)
			} else {
				// 查看栈顶算数的运算符
				topOpr := exprStack.List[exprStack.Top]
				// 获取栈顶运算符等级
				topOprLevel := exprStack.IsExprOperLevel(topOpr)
				fmt.Printf("栈顶计算符：%v \n", c)
				level := exprStack.IsExprOperLevel(ascNum)
				// 判断栈顶运算符等级是否大于待压栈的运算符，不能通过 pop 弹出，只能通过列表取值获取
				// 如果栈顶运算符 >= 准备入栈的运算符，则数栈弹出两次，进行计算; 运算符入栈
				if topOprLevel >= level {
					one, _ = numberStack.Pop()
					tow, _ = numberStack.Pop()
					opr, _ = exprStack.Pop()
					tmp = exprStack.ExprCount(one, tow, opr)
					fmt.Printf("中间计算步骤： %v %v %v = %v \n", one, opr, tow, tmp)
					numberStack.Push(tmp)
					fmt.Printf("入栈运算符是 => %v\n", c)
					exprStack.Push(ascNum)
				} else {
					fmt.Printf("栈顶运算符小于准备入栈的运算符：%v \n", c)
					exprStack.Push(ascNum)
				}
			}
		} else {
			joinNumber += c
			// 如果滑动 index == 算数表达式长度减 1 那证明已经到了算数表达式末尾
			if index == len(s)-1 {
				val, _ := strconv.Atoi(joinNumber)
				fmt.Printf("算数入栈数据: %v,算数类型 :%T\n", val, val)
				numberStack.Push(val)
				// 滑动道当前index 的后一位，判断是否是运算符，如果是运算符则证明，多位数已到最后
			} else if exprStack.IsOper(int([]byte(s[index+1 : index+2])[0])) {
				val, _ := strconv.ParseInt(joinNumber, 10, 64)
				numberStack.Push(int(val))
				joinNumber = ""
			}

		}
		// 判断当 index+1 == 算式长度的时候则需要跳出循环
		if index+1 == len(s) {
			fmt.Println("index 扫描结束，跳出第一个循环")
			break
		}

		// index 自加使得扫描位移
		index++
	}
	numberStack.Show()
	fmt.Println("数栈栈顶值：", numberStack.List[numberStack.Top])
	exprStack.Show()
	for {
		if exprStack.IsNil() {
			n, _ := numberStack.Pop()
			fmt.Printf("数栈中最后的数: %v\n", n)
			break
		}
		one, _ = numberStack.Pop()
		tow, _ = numberStack.Pop()
		opr, _ = exprStack.Pop()
		tmp = exprStack.ExprCount(one, tow, opr)
		fmt.Printf("中间计算步骤： %v %v %v = %v \n", one, opr, tow, tmp)
		numberStack.Push(tmp)

	}
	//latest, _ := numberStack.Pop()

}

func main() {
	exprNumber("3+3*5-4-9")
}
