package main

import "fmt"

func exchangeNum(num1 int, num2 int) {
	var t int
	t = num1
	num1 = num2
	num2 = t
}

func main() {
	var num1 int = 10
	var num2 int = 20
	fmt.Printf("交换前的两个数：num1=%v,num2=%v\n", num1, num2)
	exchangeNum(num1, num2)
	fmt.Printf("交换后的两个数：num1=%v,num2=%v\n", num1, num2)

}
