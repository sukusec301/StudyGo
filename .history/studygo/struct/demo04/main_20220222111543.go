package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("%p", &a)
	fmt.Printf("%p", b)
	fmt.Printf("%v", b)
	fmt.Printf("%p\n", &b)
	fmt.Printf("%v\n", &b)
}
