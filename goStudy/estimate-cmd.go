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
