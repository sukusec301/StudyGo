package main

import "fmt"

func main() {
	s1 := "beijing北京你好！"
	sum := 0
	for _, j := range s1 {
		fmt.Printf("%c\n", j)

	}
	// fmt.Println(sum)
}
