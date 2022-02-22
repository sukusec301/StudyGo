package main

import "fmt"

func main() {
	var s1 []int
	var s2 []string
	fmt.Println(s1, s2)
	//初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"beijing", "shanghai", "guangdong"}
	fmt.Println(s1, s2)
}
