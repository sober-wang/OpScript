package main

import "fmt"

type Man interface {
	name(n string) string
	age() int
}

type Woman struct {
	discript string
}


func (woman Woman) name(n string) string{
	womanName := n + ". Mei"
	woman.discript = "特蕾莎.梅 是英国首相"
	fmt.Println(woman.discript)
	return womanName
}

func (woman Woman) age() int{
	return 52
}



func main(){
	var man Man
	firstName := "Teleish"
	man = new(Woman)
	fmt.Println(man.name(firstName))
	fmt.Println(man.age())
}
