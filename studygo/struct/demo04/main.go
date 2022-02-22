package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	// var p person

	// fmt.Printf("%T\n", p)  //p的类型为main.person
	// fmt.Printf("%v\n", p)  //p的值为{0 }
	// fmt.Printf("%p\n", &p) //p的指针为0xc000004078
	var p1 = new(person)
	fmt.Printf("%T\n", p1)  //p1的类型为*main.person
	fmt.Printf("%#v\n", p1) //p1的值为&{ 0}
	fmt.Printf("%p\n", p1)  //p1的值为0xc000004078,p1保存的值就是一个内存地址
	fmt.Printf("%p\n", &p1) //p1本身就是一个变量，本身也会有自己的地址

	//结构体的初始化2：key-value初始化
	var p3 = &person{
		name: "元帅",
		age:  18,
	}
	fmt.Printf("%T\n", p3)
	//结构体的初始化3：使用值列表的形式初始化，值的顺序要和结构体定义时字段的顺序一致
	p4 := &person{
		"小王子",
		19,
	}
	fmt.Printf("%T\n", p4)
}
