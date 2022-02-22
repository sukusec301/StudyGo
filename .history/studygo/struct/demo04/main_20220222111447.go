package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("%p", &a)
	fmt.Printf("%p", b)
}
