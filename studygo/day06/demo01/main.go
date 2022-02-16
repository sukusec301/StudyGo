package main

import "fmt"

// func main() {
// 	var num1 int = 100
// 	var num2 int = 200
// 	var sum int = 0
// 	sum += num1
// 	sum += num2

// 	fmt.Println(sum)
// }
func sum(x int, y int) int {
	return x + y
}
func main() {
	// a := 10
	// b := 20
	c := sum(10, 20)
	fmt.Println(c)
}
