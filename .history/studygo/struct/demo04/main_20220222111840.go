package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var p person           //p的类型为main.person
	fmt.Printf("%p\n", &p) //p的值为{0 }
}
