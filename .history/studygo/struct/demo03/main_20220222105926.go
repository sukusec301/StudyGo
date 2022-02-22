package main

import "fmt"

type person struct {
	name, gender string
}

func changeGender(x *person) {
	(*x).gender = "BBBBBBBBBBBBBBBBBBBBBBoy"
}
func main() {
	var p person

	p.name = "holyshit"
	p.gender = "GGGGGGGGGGGGGGGGGirl"
	fmt.Println(p.gender)
	changeGender(&p)
	fmt.Println(p.gender)
	fmt.Println(p.name)
}
