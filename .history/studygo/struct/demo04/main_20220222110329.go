package main

import "fmt"

type person struct {
	name, gender string
}

func changeGender(x *person) {
	//根据内存地址找到那个变量，修改的就是原来的变量
	//	(*x).gender = "BBBBBBBBBBBBBBBBBBBBBBoy"
	x.gender = "BBBBBBBBBBBBBBBBBBBBBBoy"
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
