package main

import "fmt"

func main() {
	s1 := "How do you do"
	//先遍历出来每个字符
	//再判断是不是第一次出现
	//如果是第一次出现就直接=1
	//如果不是第一次出现，+1
	for _, j := range s1 {
		fmt.Println(j)
	}
}
