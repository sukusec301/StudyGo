package main

import "fmt"

func funcA(num1, num2 int) int {
	sum := num1 + num2
	return sum
}
func main() {
	ret := funcA(1, 2)
	fmt.Println(ret)
}
