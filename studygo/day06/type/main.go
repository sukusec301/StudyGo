package main

import "fmt"

func main() {
	type myInt int
	var num1 myInt = 30
	fmt.Println("num1", num1)

	var num2 int = 30
	num2 = int(num1) //虽然是别名，但是在go中编译是别的时候
	//仍然认为两者的数据类型不一样
	fmt.Println("num2:", num2)
	// a, b := cal(20, 10)
	// fmt.Println(a, b)
	bm, bn := cal(20, 30)
	fmt.Println(bn, bm)
}

//定义一个计算和、差的函数
func cal(x int, y int) (sum int, sub int) {

	sum = x + y
	sub = x - y
	return
}
