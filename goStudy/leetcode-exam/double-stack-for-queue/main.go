package main

import (
	"container/list"
	"errors"
	"fmt"
)

type Queue struct {
	MaxTop int
	Top    int
	list   []int
	QType  string
}

type CQueue struct {
	One    Queue
	Second Queue
}

func Constructor() CQueue {

	var cq CQueue
	cq.One = Queue{
		MaxTop: -1,
		Top:    -1,
		QType:  "one",
	}
	cq.Second = Queue{
		MaxTop: -1,
		Top:    -1,
		QType:  "second",
	}

	return cq
}

// AppendTail 插入的时候始终从一栈插入
func (this *CQueue) AppendTail(value int) {
	_ = this.One.Push(value)
	/*
		if this.One.IsFull() {
			fmt.Println("First Queue is full ,Can't append element")
			return
		}else if this.Second.Top == this.One.MaxTop {
			fmt.Println("First Queue is full ,Can't append element")
			return
		}
		// 第一种情况：第一个栈空，第二个栈非空，添加前先将二栈内的数据移至一栈
		if !this.Second.IsNil() && this.One.IsFull() {
			for i := 0; i < this.Second.Top; i++ {
				oldValue, err := this.Second.Pop()
				if err != nil {
					fmt.Println(err)
				} else {
					if err := this.One.Push(oldValue); err != nil {
						fmt.Println(err)
					}
				}
			}
			if err := this.One.Push(value); err != nil {
				fmt.Println(err)
			}
			return
		}
		// 第二种情况：二栈空，一栈有余额，数据入栈
		if this.Second.IsNil() && !this.One.IsNil() {
			this.One.Push(value)
			return
		}
	*/
}

func (this *CQueue) DeleteHead() int {
	// 先判断第二个栈是否为空
	if this.Second.IsNil() {
		// 如果为空将一栈数据转移至第二个栈
		for !this.One.IsNil() {
			// 因为是利用双栈实现队列所以双栈大小相同
			value, _ := this.One.Pop()
			//fmt.Printf("%s 栈 %d 出栈\n", this.One.QType, value)
			_ = this.Second.Push(value)
		}
	}
	if !this.Second.IsNil() {
		resultValue, _ := this.Second.Pop()
		//fmt.Printf("%s 栈 %d 出栈\n", this.Second.QType, resultValue)
		return resultValue
	}
	return -1
}

func (q *Queue) IsNil() bool {
	if q.Top == -1 {
		return true
	}
	return false
}

func (q *Queue) IsFull() bool {
	if q.MaxTop == q.Top {
		return true
	}
	return false
}

func (q *Queue) Push(value int) error {
	/* 这里注释掉为了防指 leetcode 官方案例压栈元素过多 */
	if q.MaxTop == q.Top && q.MaxTop != -1 {
		return errors.New("Stack is full")
	}
	q.Top++
	//q.list[q.Top] = value
	q.list = append(q.list, value)
	//fmt.Printf("%d 入栈. %s 栈：%v\n", value, q.QType, q.list)
	return nil
}

func (q *Queue) Pop() (int, error) {
	var value int = -1
	if q.Top == -1 {
		return value, errors.New("Stack is nil")
	}
	value = q.list[q.Top]
	fmt.Printf("栈: %v; %v 出栈; 栈内数据：%v\n", q.QType, value, q.list)
	q.list = q.list[:len(q.list)-1]
	fmt.Printf("栈：%v; %v 出栈; 栈内数据：%v\n", q.QType, value, q.list)
	q.Top--
	return value, nil

}

/* leetcode 官方解，利用 Golang 官方库 container/list */

type LeetCodeCQueue struct {
	First  *list.List
	Second *list.List
}

func LCConstructor() LeetCodeCQueue {
	return LeetCodeCQueue{
		First:  list.New(),
		Second: list.New(),
	}
}

func (lcThis *LeetCodeCQueue) AppendTail(value int) {
	lcThis.First.PushBack(value)
}

func (lcThis *LeetCodeCQueue) DeleteHead() int {
	if lcThis.Second.Len() == 0 {
		for lcThis.First.Len() > 0 {
			lcThis.Second.PushBack(lcThis.First.Remove(lcThis.First.Back()))
		}
	}

	if lcThis.Second.Len() != 0 {
		e := lcThis.Second.Back()
		lcThis.Second.Remove(e)
		return e.Value.(int)
	}
	return -1
}

/* leetcode 官方解，利用 Golang 官方库 container/list */
func main() {
	fmt.Println("vim-go")
	obj := Constructor()
	fmt.Println(obj.DeleteHead())
	obj.AppendTail(5)
	obj.AppendTail(2)

	fmt.Println("DeleteHead => ", obj.DeleteHead())
	fmt.Println("DeleteHead =>", obj.DeleteHead())
	fmt.Printf("One List => %v , Tow List=> %v\n", obj.One.list, obj.Second.list)
	///	fmt.Printf("%s 栈：%v; %s 栈：%v", obj.One.QType, obj.One.Top, obj.Second.QType, obj.Second.Top)

	lcObj := LCConstructor()
	fmt.Println(lcObj.DeleteHead())
	lcObj.AppendTail(5)
	lcObj.AppendTail(2)

	fmt.Println("DeleteHead => ", lcObj.DeleteHead())
	fmt.Println("DeleteHead =>", lcObj.DeleteHead())

}
