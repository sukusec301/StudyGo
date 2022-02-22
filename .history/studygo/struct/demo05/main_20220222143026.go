package main

import "fmt"

type person struct {
	name string
	age  int64
}

func main() {
	//结构体初始化方式一：
	var p1 person
	p1.name = "第一种初始化方式"
	p1.age = 18

	//结构体初始化方式二：
	p2 := person{
		name: "第二种初始化方式",
		age:  19,
	}

	//结构体初始化方式三：
	p3 := person{
		"第三种初始化方式",
		20,
	}

	fmt.Printf("p1为%v\n", p1)
	fmt.Printf("p2为%v\n", p2)
	fmt.Printf("p3为%v\n", p3)
}
