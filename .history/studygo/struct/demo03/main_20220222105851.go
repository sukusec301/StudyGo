package main

import "fmt"

type person struct {
	name, gender string
}

func changeGender(x *person) (y *person) {
	(*x).gender = "BBBBBBBBBBBBBBBBBBBBBBoy"
	return
}
func main() {
	var p person

	p.name = "holyshit"
	p.gender = "GGGGGGGGGGGGGGGGGirl"
	fmt.Println(p.gender)
	ret := changeGender(&p)
	fmt.Println(ret)
	fmt.Println(p.name)
}
