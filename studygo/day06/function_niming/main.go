package main

import "fmt"

//在函数外部新建一个最简单的匿名函数，不过这样子并不推荐
var a = func() {
	fmt.Println("hjellow!")
}

func main() {
	a()
	//函数内部调用匿名函数，相当于用一个变量声明匿名函数，随后调用
	f1 := func() {
		fmt.Println("helkad")
	}
	f1()
	//如果只是调用一次，可以简写称立即执行函数
	func() {
		fmt.Println("立即执行匿名函数")
	}() //这里就是立即执行

	//立即调用带有参数的匿名函数
	func(x int, y int) {
		ret := x + y
		fmt.Println(ret)
	}(10, 20)
}
