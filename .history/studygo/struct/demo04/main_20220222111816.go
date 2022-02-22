package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var p person //p的类型为main.person
	fmt.Printf("%v\n", p)
}
