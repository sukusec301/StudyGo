package main

import "fmt"

//变量的作用域

//定义一个全局变量
func f1() {
	var x int = 100
	fmt.Println(x)

}

func main() {
	iff()
	fuck()
	f1()
}
func iff() {
	if y := 35; y > 19 {
		fmt.Println("你好")
	}
}
func fuck() {
	for i := 10; i < 100; i++ {
		fmt.Println(i)
	}
}
