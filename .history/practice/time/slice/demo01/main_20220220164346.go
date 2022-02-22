package main

import "fmt"

func main() {
	//切片的定义方式1：直接定义
	var slice1 []int
	slice1 = []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)

	//切片的定义方式2：从数组中取出切片
	a1 := [...]int{1, 2, 3, 3, 4, 12, 3}
	s3 := a1[0:4]
	fmt.Println(s3)

	s4 := a1[:4]
	fmt.Println(s4)

}
