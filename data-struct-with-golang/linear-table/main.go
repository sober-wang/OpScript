package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

const maxSize int = 10

type linearNode struct {
	value  [maxSize]string // 线性表的值
	length int             // 线性表节点当前位置
}

// NeNewLinearTable 初始化一个线性表
func NewLinearTable() linearNode {
	var ln linearNode
	ln.length = 0
	return ln
}

// linearInsert 往线性表里面插入数据
func (l *linearNode) linearInsert(site int, value string) error {
	// 定义一个错误
	s := fmt.Sprintf("插入 [ %v ] 位置不合法", site)
	switch {
	// 判断线性表是否已满
	case l.length == maxSize:
		return errors.New("线性表已满")
		// 由于是用数组实现，数组线标不能越界，否则会 panic
	case site < 0 || site > maxSize-1:
		return errors.New(s)
		// site 的位置不能大于现在线性表记录的长度，并且不能超过线性表大小
	case site > l.length+1 && site < maxSize-1:
		return errors.New(s)
	default:
		// 在插入的目标节点把对应的数据向后移位
		for i := l.length; i >= site; i-- {
			l.value[i+1] = l.value[i]
		}
		l.value[site] = value
		// 线性表大小 +1
		l.length++
		return nil
	}
}

// linearDelete 删除线性表元素
func (l *linearNode) linearDelete(site int) {
	if l.length == 0 {
		fmt.Println("线性表为空不能删除")
		return
	}
	// 需要删除的位置不能小于等于 0, 我们生活中的案例常从 1 开始计算表大小 ，或者大于线性表现在的长度
	if site <= 0 || site > l.length-1 {
		fmt.Printf("删除元素位置不正确, 删除位置 小于线性表长度 site:= [ %v ] ,length := [ %v ]\n", site, l.length)
		return
	}
	// 删除元素，线标应该从传入指定下标 - 1 开始，把后面的数据向前挪移。
	for i := site - 1; i < l.length; i++ {
		l.value[i] = l.value[i+1]
	}
	// 要把最后一个位置的的数据清空
	l.value[l.length-1] = ""
	fmt.Printf("你删除了第 [ %v ] 位元素\n", site)
	// 线性表长度减少一个
	l.length--
	return
}

func (l *linearNode) find(site int) bool {
	if l.length == 0 || site < 0 || site > l.length {
		return false
	}
	fmt.Println(l.value[site-1])
	return true
}

func main() {
	lt := NewLinearTable()
	data, _ := json.Marshal(lt)
	fmt.Println("初始化 => ", string(data))
	_ = lt.linearInsert(0, "a")
	_ = lt.linearInsert(1, "b")
	_ = lt.linearInsert(2, "c")
	if err := lt.linearInsert(5, "mxcivo"); err != nil {
		fmt.Println(err)
	}
	if err := lt.linearInsert(100, "jsiodjfojnkl34"); err != nil {
		fmt.Println(err)
	}

	_ = lt.linearInsert(3, "z")
	fmt.Println(lt)

	lt.find(3)
	lt.linearDelete(2)
	fmt.Println(lt)
}
