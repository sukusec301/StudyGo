package main

import "fmt"

func f1() {
	fmt.Println("和姥姥的骄傲大家")
	fmt.Println("和姥姥的骄傲大家")
	fmt.Println("和姥姥的骄傲大家")
}
func f2(x string) {
	fmt.Println("Hello!", x)
}

//参数类型简写
func f3(x, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return
}

//可变参数（多个参数）
func f4(x int, y ...string) {
	fmt.Println(x, y)
}

func main() {
	f1()
	f2("北京")
	f2("上海")
	f3(100, 200)
}

//函数中不支持再添加一个函数，匿名函数可以
