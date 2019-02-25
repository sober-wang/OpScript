package main

/*
	JSON 创建学习，嵌套多层 JSON 数据
*/

import (
	"encoding/json"
	"fmt"
)

// 第二层 JSON 的 struct 创建
type Hello struct {
	CubeType string
	High     int
	Width    int
	Long     int
}

// 第一层 JSON 的 struct 创建
type Study struct {
	Name      string
	Age       int
	Skill     []string
	SkillType map[string]string
	// 在 A struct 中使用 B struct 就需要指定 B 的内存地址
	Cube *Hello
}

func main() {
	// 创建指针将 st 的地址指向 Study struct 地址,并给变量赋值
	st := &Study{
		"sober",
		26,
		[]string{"Hadoop", "Linux", "Python", "Golang"},
		map[string]string{
			"programa": "Python",
			"OS":       "Linux",
		},
		// 创建第二层 JSON 的指针，并给变量赋值
		&Hello{
			CubeType: "三角型",
			High:     1,
			Width:    1,
			Long:     1,
		},
	}

	// 将数据编码为 JSON 数据
	b, err := json.Marshal(st)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
	fmt.Printf("this is %s\n", b)
	//stb := &Study{}
}
