package main

type person struct {
	name int64
	age  int64
}

func main() {
	//结构体初始化方式一：
	var p1 person
	p1.name = "第一种初始化方式"
	p1.age = "第一种初始化方式"
	p1.gender = "第一种初始化方式"
	p1.name = "第一种初始化方式"
}
