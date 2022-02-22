package main

import "fmt"

func funcA() {
	fmt.Println("func A")
}

func funcB() {
	panic("出错啦！！")
}

func funcC() {
	fmt.Println("func C")
}
func main() {
	funcA()
	funcB()
	funcC()
}
