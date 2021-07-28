package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type ListNode struct {
	Next  *ListNode
	Value int
}

func reversePrint(head *ListNode) []int {
	var l []int
	newhead := &ListNode{}
	// 定义一个无大小的栈
	s := Stack{
		MaxTop: -1,
		Top:    -1,
		List:   newhead,
	}
	tmp := head
	if tmp.Next == nil {
		return l
	}
	for {
		tmp = tmp.Next
		// 用来存储顺序输出的列表
		l = append(l, tmp.Value)
		s.Push(tmp.Value)
		if tmp.Next == nil {
			//l = append(l, tmp.Value)
			break
		}
	}
	fmt.Println("顺序输出:", l)
	fmt.Println("栈大小：", s.Top)

	var rl []int
	for {
		val, err := s.Pop()
		//fmt.Println("倒叙输出: ", val)
		if err != nil {
			break
		}
		rl = append(rl, val)
		//rl[i] = val
		//	fmt.Println(n)

	}
	fmt.Println("倒叙输出:", rl)

	return rl

}

func AddTailNode(head *ListNode, value int) {
	tmp := head
	for {
		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}
	var newNode ListNode
	newNode.Value = value
	tmp.Next = &newNode
	return
}

// AddHeadNode 由于要利用栈，所以链表要从头部插入；完全不使用 go 自带的容器例如: 切片，数组
func (ln *ListNode) AddHeadNode(value int) {
	node := &ListNode{
		Value: value,
		Next:  ln.Next,
	}
	ln.Next = node
}

// RemoveHeadNode 基于链表的栈实现，所以需要从链表头部删除。
func (ln *ListNode) RemoveHeadNode() int {
	tmp := ln
	if tmp.Next == nil {
		return 0
	}
	rv := tmp.Next.Value
	ln.Next = tmp.Next.Next
	return rv
}

func RemoveTailNode(head *ListNode) int {
	tmp := head
	if tmp.Next == nil {
		//head.Value = 0
		return tmp.Value
	}
	for {
		if tmp.Next.Next == nil {
			break
		}
		tmp = tmp.Next
	}
	var rv int
	rv = tmp.Next.Value
	tmp.Next = nil
	return rv

}

func Show(head *ListNode) {
	tmp := head
	if tmp.Next == nil {
		return
	}
	for {

		tmp = tmp.Next
		if tmp.Next == nil {
			break
		}
	}
}

type Stack struct {
	MaxTop int
	Top    int
	List   *ListNode
}

func (s *Stack) Push(value int) {
	if s.MaxTop == s.Top && s.MaxTop != -1 {
		return
	}
	//s.List.Next.Value = value
	s.List.AddHeadNode(value)
	s.Top++
}

func (s *Stack) Pop() (int, error) {
	if s.Top == -1 {
		return 0, errors.New("error")
	}

	var rv int
	rv = s.List.RemoveHeadNode()
	//fmt.Printf("%v 出栈\n", rv)
	//fmt.Println("Top Value => ", s.Top)
	s.Top--
	return rv, nil
}

func main() {
	head := &ListNode{}
	/*
		AddTailNode(head, 1)
		AddTailNode(head, 3)
		AddTailNode(head, 2)
					AddTailNode(head, 3)
					AddTailNode(head, 2)
					AddTailNode(head, 2)
					AddTailNode(head, 1)
					AddTailNode(head, 2)
					AddTailNode(head, 1)
					AddTailNode(head, 3)
					AddTailNode(head, 2)
					AddTailNode(head, 2)
					AddTailNode(head, 2)
					AddTailNode(head, 1)
					AddTailNode(head, 3)
	*/
	for i := 0; i < 1000; i++ {
		val := rand.Intn(9)
		AddTailNode(head, val)
		//head.AddHeadNode(val)
	}

	reversePrint(head)
}
