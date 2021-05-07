package main

import (
	"fmt"
	"sync"
)

type single struct{}

var hungerSingle *single = &single{}

// hungerSingle 饥饿模式
func hgSingle() *single {
	fmt.Println("创建实例")
	return hungerSingle
}

// lazySingle()
func lazyUnsafeSingle() {
	var lus *single
	lusFunc := func() {
		if lus == nil {
			fmt.Println("创建实例")
			lus = &single{}
		}
		return
	}

	// 启动多个 goroutine 创建 single
	for i := 0; i < 10000; i++ {
		go lusFunc()
	}

}

func lazySafeSingle() {
	var lus *single
	var m sync.Mutex
	lusFunc := func() {
		m.Lock()
		if lus == nil {
			fmt.Println("创建实例")
			lus = &single{}
		}
		m.Unlock()
		return
	}

	// 启动多个 goroutine 创建 single
	for i := 0; i < 10000; i++ {
		go lusFunc()
	}

}

func goExclusiveSingle() {
	var once sync.Once
	var inst *single
	gels := func() {
		once.Do(func() {
			fmt.Println("创建实例")
			inst = &single{}
		})
	}

	// 启动多个 goroutine 创建 single
	for i := 0; i < 10000; i++ {
		go gels()
	}

}

func main() {
	fmt.Println("vim-go")
	hgSingle()
	lazyUnsafeSingle()
	fmt.Println("------------------懒汉安全模式开始创建--------------------")
	lazySafeSingle()
	fmt.Println("------------------Golang 专属模式创建--------------------")
	goExclusiveSingle()
}
