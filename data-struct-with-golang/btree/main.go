package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	Number int
	Left   *Node
	Right  *Node
}

func MakeBtree() *Node {
	return &Node{}
}

func (n *Node) IsNil() bool {
	if n == nil {
		return true
	}
	return false
}

func (n *Node) AddNode(d *Node) {

	if n.IsNil() {
		n.Number = d.Number
		return
	}
	// 当新节点的ID > 当前节点ID 就添加在右侧
	if d.Number > n.Number {
		// 如果当前节点右侧为空,则将新节点赋值给当前节点的右子节点
		if n.Right == nil {
			n.Right = d
		} else {
			// 当前节点右子结点不为空，递归调用当前方法
			n.Right.AddNode(d)
		}
		// 当新节点ID < 当前节点ID，就添加在左侧
	} else {
		// 当左子结点为空就将新节点赋值给左子节点
		if n.Left == nil {
			n.Left = d
		} else {
			// 当前节点左子结点不为空，递归调用当前方法。
			n.Left.AddNode(d)
		}
	}
}

func (n *Node) TreeFind(id int) {
	if n.IsNil() {
		fmt.Println("空树")
		return
	}
	if id == n.Number {
		fmt.Printf("找到了: %v \n", n.Number)
		return
	}

	if id > n.Number {
		n.Right.TreeFind(id)
	} else {
		n.Left.TreeFind(id)
	}
}

func (n *Node) Hight() int {
	if n.IsNil() {
		return 0
	}
	rightH := n.Right.Hight()
	leftH := n.Left.Hight()
	if rightH > leftH {
		return rightH + 1
	} else {
		return leftH + 1
	}

}

// 前序遍历
func EachBylast(n *Node) {
	if !n.IsNil() {
		fmt.Printf("节点ID： %v \n", n.Number)
		EachBylast(n.Left)
		EachBylast(n.Right)
	}
}

// createBtree()
func createBtree(root *Node) {
	for i := 1; i < 10; i++ {
		n := rand.Intn(500)
		node := &Node{
			Number: n,
		}
		root.AddNode(node)
	}
}

func main() {
	fmt.Println("vim-go")
	head := MakeBtree()
	createBtree(head)
	EachBylast(head)
	head.TreeFind(5)
	btreeHight := head.Hight()
	fmt.Printf("树的高度：%v\n", btreeHight)

}
