package main

/*
	JSON 数据解析
*/

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

// 创建一个错误处理函数，避免过多的 if err != nil{} 出现
func dropErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Hello,today we will study Golang read JSON data")

	// 获取 参数，请传入一个文件路径
	filePath := os.Args[1]
	fmt.Printf("The file path is :%s\n", filePath)

	// ioutil 方式读取，会一次性读取整个文件，在对大文件处理时会有内存压力
	fileData, err := ioutil.ReadFile(filePath)
	dropErr(err)
	fmt.Println(string(fileData))

	// bufio 读取
	f, err := os.Open(filePath)
	dropErr(err)
	bio := bufio.NewReader(f)
	// ReadLine() 方法一次尝试读取一行，如果过默认缓存值就会报错。默认遇见'\n'换行符会返回值。isPrefix 在查找到行尾标记后返回 false
	bfRead, isPrefix, err := bio.ReadLine()
	dropErr(err)
	fmt.Printf("This mess is  [ %q ] [%v]\n", bfRead, isPrefix)

	// 解析 JSON 数据使用 json.Unmarshal([]byte(JSON_DATA),JSON对应的结构体) ,也就是说我们在解析 JSON 的时候需要确定 JSON 的数据结构
	res := &Study{}
	json.Unmarshal([]byte(bfRead), &res)
	fmt.Println(res.Cube)

}
