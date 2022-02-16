package main

import "fmt"

func test(num *int) {
	*num = 30
}

func main() {
	var num int = 10
	fmt.Println(&num)
	test(&num)
	fmt.Println(num)
}
