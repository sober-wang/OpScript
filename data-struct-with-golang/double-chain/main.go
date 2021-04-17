package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	Number   int
	Next     *Node
	Previous *Node
}

func isNil(head *Node) bool {
	if head == nil {
		return true
	}
	return false
}

func showNode(head *Node) {
	tmp := head
	if isNil(tmp) {
		fmt.Println("双向链表为空")
		return
	}
	for {
		if tmp.Next == nil {
			fmt.Println("双向链表完成打印")
			break
		}
		fmt.Printf("当前节点ID: [ %v ]\n", tmp.Number)
		tmp = tmp.Next
	}
}

// delNode 删除目标结点
func delNode(head *Node, nv int) {
	tmp := head
	check := false
	if isNil(tmp) {
		fmt.Println("双向链表为空")
		return
	}
	for {
		// 因为是 节点从 0 开始所以 nv-1 才能删除目标节点
		if tmp.Next == nil && tmp.Number != nv-1 {
			break
		} else if tmp.Number == nv-1 {
			check = true
			break
		}
		tmp = tmp.Next
	}

	if check {
		tmp.Next = tmp.Next.Next
		// 如果 tmp.Next 不为空，证明结点后还有结点，那么就需要将下一个结点的前指针指向 tmp
		if tmp.Next != nil {
			tmp.Next.Previous = tmp
		}
	} else {
		fmt.Printf("不存在该节点ID：%v\n", nv)
	}

}

// flashBack 倒叙输出双向链表
func flashBack(head *Node) {
	tmp := head
	if isNil(tmp) {
		fmt.Println("双向链表为空")
		return
	}

	// tmp 先到滑动到尾部
	for {
		if tmp.Next == nil {
			fmt.Println("滑动节点已到尾部，准备倒叙输出")
			break
		}
		tmp = tmp.Next
	}

	// 开始倒叙输出
	for {
		if tmp.Previous == nil {
			fmt.Println("倒叙输出已完成")
			break
		}
		fmt.Printf("当前节点：[ %v ]\n", tmp.Number)
		tmp = tmp.Previous
	}
}

func insertNode(head *Node, node *Node) {
	tmp := head
	if isNil(tmp) {
		fmt.Println("双向链表为空")
		return
	}

	for {
		if tmp.Next == nil {
			//tailAddNode(tmp, node)
			tmp.Next = node
			//fmt.Println("tmp.Next == nil已到链尾")
			break
			// 顺序插入结点
		} else if tmp.Number < node.Number && tmp.Next.Number > node.Number {
			node.Next = tmp.Next
			tmp.Next.Previous = node
			node.Previous = tmp
			tmp.Next = node
			break
		}
		tmp = tmp.Next
	}
}

// tailAddNode 尾部添加结点
func tailAddNode(head *Node, node *Node) {
	tmp := head
	if isNil(tmp) {
		fmt.Println("双向链表为空")
		return
	}
	for {
		if tmp.Next == nil {
			tmp.Next = node
			node.Previous = tmp
			break
		}
		tmp = tmp.Next
	}

}

// batchDoubleChain 批量添加双向链表结点
func batchDoubleChain(head *Node, num int) {
	for i := 1; i <= num; i++ {
		node := &Node{
			Number: i,
		}
		insertNode(head, node)
	}
}

/*
func sortRandNode(head *Node){
	newHead = &Node{}
	tmp := head

	insert := func(newh *Node,node *Node){

	}

	show := func(h *Node){
		t := h
		if isNil(t){
			fmt.Println("双向链表为空")
			return
		}
		for {
			if t.Next == nil{

			}
		}
	}

}
*/

func batchRandNode(head *Node, num, bitSize int) {
	for i := 1; i < num; i++ {
		node := &Node{
			Number: rand.Intn(bitSize),
		}
		tailAddNode(head, node)

	}
}

func main() {
	fmt.Println("vim-go")

	head := &Node{}
	/*
			first := &Node{
				Number: 223,
			}
			second := &Node{
				Number: 1123,
			}
			tailAddNode(head, first)
			tailAddNode(head, second)
		iNode := &Node{
			Number: 225,
		}
		insertNode(head, iNode)
	*/
	batchDoubleChain(head, 10)
	//batchRandNode(head, 50, 500)
	delNode(head, 3)
	showNode(head)
	//flashBack(head)

}
