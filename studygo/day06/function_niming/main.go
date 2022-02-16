package main

import "fmt"

//在函数外部新建一个最简单的匿名函数，不过这样子并不推荐
var a = func() {
	fmt.Println("hjellow!")
}

func main() {
	a()
}
