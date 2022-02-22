package main

import "fmt"

func main() {
	s1 := "beijing北京你好！"
	sum := 0
	for i := 0; i < len(s1); i++ {
		if len(s1[i]) == 3 {
			sum++
		}
	}
	fmt.Println(sum)
}
