package main

import "fmt"

type person struct {
	name string
}

func main() {
	var p person
	p.name = "hello!!"
	fmt.Println(&p.name)
	fmt.Println(&p)
}
