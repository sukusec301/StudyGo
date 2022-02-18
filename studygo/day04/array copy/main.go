package main

import "fmt"

func testarray() {
	array := [...]int{1, 2, 312, 31, 23, 124, 1, 451, 45}
	sum := 0
	for _, j := range array {
		sum := sum + j
		fmt.Println(sum)
	}
}
func main() {
	testarray()
	sum := arraySum([3]int{1, 2, 3})
	fmt.Println(sum)
}
func arraySum(x [3]int) int {
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}
