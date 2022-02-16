package main

import "fmt"

func f1() {
	fmt.Println("Hellow shahe")
}
func f2() int {
	// fmt.Println("hellow")
	return 10
}

func f3(y func()) {
	fmt.Println("asd")
}
func main() {
	a := f1
	b := f2
	fmt.Printf("%T\n%T", a, b)
	f3(f1)
	f5()
	// f3(b)
}

//函数还可以作为返回值
func f5(x func() int) func(int, int) int {
	return ff
}
func ff(a, b int) int {
	return a + b
}
