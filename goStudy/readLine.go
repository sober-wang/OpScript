/*
	os.Args[1] 可以获取命令行参数，os.Args[0] 是执行文件本体
	defer 在函数执行完成后做的操作，防止忘记关闭链接，关闭文件等，或者在函数执行完成后去做一些其他的操作。比较消耗资源
	time.Sleep() 默认秒数为纳秒，非常小的一个单位。通常配合 time 中的其他函数使用
	bufio 库可以实现对文件的操作。在bufio库中可以设置数据读写缓存设置，这样可以对读写进行加速

*/
package main

import (
	"io"
	"fmt"
	"bufio"
	"os"
	"time"
	"strings"
)

func main(){
	fileName := os.Args[1]
	fmt.Println("The file name is ",fileName)
	f,err := os.Open(fileName)
	if err != nil{
		fmt.Println("This have a ERROR",err)
		return
	}
	defer f.close()

	br := bufio.NewReader(f)
	for {
		s,_,c := br.ReadLine()
		if c == io.EOF {
			break
		}
		time.Sleep(100000000)
		fmt.Printf("s type: %T \n",s)
		fmt.Println(cutString(string(s)))
	}
}

func cutString(msg string) []string {
	newString := strings.SplitN(msg," ",-1)
	return newString
}
