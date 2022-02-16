package main

import "fmt"

func cal(num1 int, num2 int) (int, int) {
	var sum int = 0
	var sub int = 0
	sum = num1 + num2
	sub = num1 - num2
	return sum, sub
	// fmt.Println()
}

func main() {
	sum1, sub1 := cal(10, 20)
	//sum1, _ := cal(10, 20)
	fmt.Println(sum1, sub1)
}
