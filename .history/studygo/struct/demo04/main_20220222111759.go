package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var p person
	fmt.Printf("%T\n", p)
}
