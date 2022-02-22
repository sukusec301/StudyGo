package main

import "fmt"

type person struct {
	name string
}

func main() {
	var p person
	fmt.Println(&p)
}
