package main

import "fmt"

type Node struct {
	Number int
	Next   *Node
}

// tailAddNode 从尾部添加节点
func tailAddNode(head *Node, newNode *Node) {
	tmp := head
	for {
		if tmp.Next == nil {
			tmp.Next = newNode
			break
		}
		tmp = tmp.Next
	}
}

// showNode 打印整个链表
func showNode(head *Node) {
	tmp := head
	for {
		if tmp.Next == nil {
			//fmt.Printf("当前节点: [ %v ]\n", tmp.Number)
			fmt.Println("单向链表打印完成")
			break
		}
		tmp = tmp.Next
		// 这里要注意在 go 语言中 int 变量初始值为 0 ,如果在 tmp 滑动前打印会把 head 节点
		fmt.Printf("当前节点: [ %v ]\n", tmp.Number)
	}
}

// addABatch 批量创建链表中的元素
func addABatch(head *Node, num int) {
	for i := 0; i < num; i++ {
		node := &Node{
			Number: i,
		}
		tailAddNode(head, node)
	}
}

func isNil(head *Node) bool {
	if head == nil {
		return true
	}
	return false
}

// insertNode 顺序插入节点
func insertNode(head *Node, node *Node) {
	tmp := head
	if isNil(tmp) {
		fmt.Println("链表为空")
		return
	}

	for {
		if tmp.Next == nil {
			break
		} else if tmp.Number == node.Number {
			fmt.Printf("新节点节点ID冲突不能插入")
			break
			// 插入的节点一定是大于 tmp.Number 并且小于 tmp.Nex.Number
		} else if tmp.Number < node.Number && tmp.Next.Number > node.Number {
			// 链表在插入的时候会断开，所以先找一个临时变量将后半部分存储下来
			newTmp := tmp.Next
			// 新节点的下一个节点指向链表后半部分
			node.Next = newTmp
			// 将 tmp.Next 指向新节点
			tmp.Next = node
			fmt.Printf("节点 [ %v ] 插入成功\n", node.Number)
			break
		} else {
			tmp = tmp.Next
		}
	}

}

// delNode 删除指定节点
func delNode(head *Node, id int) {
	// 定义滑动变量
	tmp := head
	check := false
	if isNil(tmp) {
		fmt.Println("链表为空")
		return
	}
	for {
		if tmp.Next == nil {
			fmt.Printf("已到链表尾部，未找到目标节点： [ %v ]\n", id)
		}
		// 滑动到需要删除的节点的前一个结点，tmp.Next.Number == id
		if tmp.Next.Number == id {
			check = true
			break
		}
		tmp = tmp.Next
	}

	if check {
		// 将 tmp 下一个结点指向下下个结点
		tmp.Next = tmp.Next.Next
	} else {
		fmt.Printf("未找到 [ %v ] 节点", id)
	}
}

func main() {
	fmt.Println("vim-go")
	n := 5
	head := &Node{}
	addABatch(head, n)
	newNode := &Node{
		Number: 11,
	}
	tailAddNode(head, newNode)
	tenNode := &Node{
		Number: 8,
	}
	insertNode(head, tenNode)

	ttNode := &Node{
		Number: 7,
	}
	insertNode(head, ttNode)
	nNode := &Node{
		Number: 9,
	}
	insertNode(head, nNode)
	fmt.Printf("未删除 8 号节点")
	showNode(head)
	delNode(head, 8)
	fmt.Print("已删除 8 号节点")
	showNode(head)
}
