package main

import "fmt"

func deferdemo() {
	defer fmt.Println("哈哈哈哈哈")
	defer fmt.Println("xxxxxxxxx")
	defer fmt.Println("aaaaaaaaa")
	fmt.Println("不不不不不")
	fmt.Println("对对对对对")
}
func main() {
	deferdemo()
}
func 