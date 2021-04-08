package main

import (
	"fmt"
	"log"
	"math/bits"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// findRepeatNumber 哈希表寻找重复的数字
func findRepeatNumber(nums []int) {
	nu := make(map[int]bool)
	for _, v := range nums {
		// 第一次遇见 v 肯定是不存在的 false ,那么就设置 true 作为已存在的标是。
		// 第二次遇见 v != false 了，那么就打印它！
		if !nu[v] {
			nu[v] = true
		} else {
			fmt.Println("重复的数字是 => ", v)
		}

	}

}

func replaceSpace(s string) {
	// 这里创建拼接字符串，但是需要注意每次拼接都会重建一个字符串对象会比较浪费对象，建议使用切片，然后用 strings.Join()方式合并输出
	var str string
	// var resultS []string

	for _, v := range []byte(s) {
		//fmt.Println(v)
		if string(v) == " " {
			//resultS = append(resultS, "%20")
			str += "%20"
		} else {
			//resultS = append(resultS, string(v))
			str += string(v)
		}

	}
	fmt.Printf("替换字符串: %v 中的空格为，替换后的结果: %v\n", s, str)
	//fmt.Printf("替换字符串: %v 中的空格为，替换后的结果: %v\n", s, strings.Join(resultS,""))
}

// reversePrint 倒叙输出链表内的值，
// 先遍历链表将所有的值收集出来
// 在通过遍历切片方式将值的列表倒叙输出
func reversePrint(head *ListNode) {
	tmp := head
	//var tail ListNode
	/*
		if tmp.Next == nil {
			fmt.Printf("struct: %v ; valuse: %v \n", head, head.Val)
			return
		}
		// 这里使用了递归
		if head != nil {
			reversePrint(head.Next)
			fmt.Printf("struct: %v ; valuse: %v \n", head, head.Val)
		}
	*/
	var intList []int
	for {
		if tmp.Next == nil {
			intList = append(intList, tmp.Val)
			//fmt.Println("到尾了 => ", tmp.Val)
			break
			//tail.Next = &tmp
		}
		intList = append(intList, tmp.Val)
		//fmt.Println(tmp.Val)
		tmp = tmp.Next

	}
	fmt.Printf("正序输出的值: %v\n", intList)
	var returnList []int
	for i := len(intList); i > 0; i-- {
		//fmt.Println(intList[i-1])
		returnList = append(returnList, intList[i-1])
	}
	fmt.Printf("倒序输出值: %v\n", returnList)

}

// reverseList 反转链表
func reverseList(head *ListNode) *ListNode {
	tmp := head
	var prev *ListNode
	for tmp != nil {
		// 中间变量保存下一个节点
		next := tmp.Next
		// 将头节点的next 指向 prev ，prev 当前为空
		tmp.Next = prev
		prev = tmp
		tmp = next
	}
	fmt.Println(prev)
	return prev
}

func showList(head *ListNode) {
	tmp := head
	fmt.Println("开始输出链表")
	for {
		if tmp.Next == nil {
			fmt.Println(tmp.Val)
			fmt.Println("结束输出链表")
			break
		}
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}
}

// minArray 找出数组中最小的数
func minArray(numbers []int) int {
	var tmp int
	tmp = numbers[0]
	log.Printf("初始化 tmp 值: %v\n", tmp)
	for _, v := range numbers[1:] {
		if v < tmp {
			tmp = v
		} else {
			log.Printf("当前循环的数字: %v\n", v)
		}
	}

	//log.Println(tmp)
	return tmp
}

// hammingWeight 找到二进制中 1 的数量
func hammingWeight(num uint32) {
	n := bits.OnesCount32(num)
	fmt.Printf("二进制中1的数量: %v\n", n)
}

func main() {
	str := "Hello, playground"
	a := []int{1, 3, 4, 5, 2, 53, 0, 1}
	findRepeatNumber(a)
	replaceSpace(str)

	var head ListNode

	var firHead ListNode
	head.Next = &firHead
	firHead.Val = 1
	var secHead ListNode
	firHead.Next = &secHead
	secHead.Val = 2
	var threHead ListNode
	secHead.Next = &threHead
	threHead.Val = 3
	reversePrint(&head)
	reverseListResult := reverseList(&firHead)
	showList(reverseListResult)

	minArray(a)

	var n uint32 = 00000000000000000000000000101011
	fmt.Printf("入参类型:%T \n", n)
	hammingWeight(n)

}
