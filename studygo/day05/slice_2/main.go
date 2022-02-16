package main

import "fmt"

func main() {
	// 切片的定义，其实就是[]中没有任何的数字提示，以及后面没有各种数值
	var s1 []int
	var s2 []string
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	//初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"beijing", "haidian", "niubi"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	//长度和容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))

}
