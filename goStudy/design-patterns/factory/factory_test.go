package factory

import "testing"

func TestNewCateen(t *testing.T){
	bone := "x"
	NewCateen("tang").MakeFood(bone)

	mainCourse := "meats"
	NewCateen("cai").MakeFood(mainCourse)
}