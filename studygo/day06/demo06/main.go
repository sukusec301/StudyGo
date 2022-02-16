package main

import "fmt"

func test(num int) int {
	num = 30
	return num
}

func main() {
	num1 := test(10)
	fmt.Println(num1)
}
