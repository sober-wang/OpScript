package main

import "fmt"

type menMsg struct {
	ID int
	Name string
	Age int
	Adderss string
	Salary int 
}

func main(){
	var sober menMsg
	sober.ID = 9527
	sober.Name = "Sober.Wang"
	sober.Salary -= 5000
	fmt.Println("这个人的工资现在是多少",sober.Salary)
	fmt.Println("The worker name is ",sober.Name)

	fmt.Println(sober)
}
