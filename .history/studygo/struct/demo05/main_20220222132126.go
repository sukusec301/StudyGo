package main

import "fmt"

type person struct {
	name   int64
	age    int64
	gender string
	gendeR string
}

func main() {
	var p1 person
	p1.name = 16
	p1.age = 18
	p1.gender = "14"
	p1.gendeR = "15"
	fmt.Printf("%p\n", &p1.name)
	fmt.Printf("%p\n", &p1.age)
	fmt.Printf("%p\n", &p1.gender)
	fmt.Printf("%p\n", &p1.gendeR)
	// fmt.Printf("%p\n",&p1.name)
}
