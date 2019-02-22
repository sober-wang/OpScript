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
