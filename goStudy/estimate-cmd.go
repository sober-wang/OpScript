/*
	这些代码中的知识点：
	string 是不可变字节序列，可以包含任意数据，所以底层表示为 byte类型 的数组
	unicode 中的 unicode.IsNumber() 等方法只能接收单个字符，也就是 rune 类型 或者 i := '1' 定义的数据
	通过 for 遍历 string 就可以得到 单个字符 rune 类型的数据
*/
package main

import (
	"fmt"
	"os"
	"unicode"
)

func lookCli(i string) string{

	if len(i) > 1{
		for _,v := range i {
			switch {
			case unicode.IsNumber(v) == true :
				fmt.Printf("这是一个数字：%c\n",v)
			case unicode.IsLetter(v) == true :
				fmt.Printf("这是一个字母：%c\n",v)
			default:
				fmt.Printf("这是一个特殊字符：%c\n",v)
			}
		}
		return "我已识别完"
	}else{
		return "只有一个字符"

	}
}

func main(){
	var Cli string
	Cli = os.Args[1]
	fmt.Printf("命令行传入参数为 ： %s,类型：%T\n",Cli,Cli)
	var result string
	result = lookCli(Cli)
	fmt.Println(result)
}
