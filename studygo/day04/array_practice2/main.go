package main

import "fmt"

func main() {
	//第一种写法
	sum := 0
	a1 := [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(a1); i++ {
		// fmt.Println(a1[i])
		sum = sum + a1[i]
	}
	fmt.Println(sum)

	fmt.Println("--------------------------------")
	fmt.Println("分界线======================分界线")
	fmt.Println("--------------------------------")
	//第二种写法
	sum2 := 0
	a2 := [...]int{10, 2, 30, 1, 9, 10, 11}
	for _, v := range a2 {
		sum2 = sum2 + v

	}
	fmt.Println(sum2)
}
