package main

import "fmt"

type person struct {
	name   int8
	age    int8
	gender int8
}

func main() {
	var p1 person
	p1.name = "nihao"
	p1.age = 18
	p1.gender = "boy"
	fmt.Printf("%p\n", &p1.name)
	fmt.Printf("%p\n", &p1.age)
	fmt.Printf("%p\n", &p1.gender)
	// fmt.Printf("%p\n",&p1.name)
}
