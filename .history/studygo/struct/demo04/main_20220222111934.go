package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var p person

	fmt.Printf("%T\n", &p) //p的类型为main.person
	fmt.Printf("%v\n", &p) //p的值为{0 }
	fmt.Printf("%p\n", &p) //p的值为{0 }
}
