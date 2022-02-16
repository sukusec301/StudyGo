package main

import (
	"fmt"
)

func test(a int) {
	fmt.Println(a)
}
func main() {

	b := test
	//函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了
	fmt.Printf("test函数的类型为：%T\nb的类型为：%T", test, b)
	//通过该变量可以对函数调用
	b(10)
}
