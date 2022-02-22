package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", b)
	fmt.Printf("%v\n", b)
	fmt.Printf("%p\n", &b)
	fmt.Printf("%v\n", &b)
}
