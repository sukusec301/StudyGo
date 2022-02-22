package main

import (
	"fmt"
	"unicode"
)

func main() {
	s1 := "beijing你好啊愛してる "
	sum := 0
	for _, j := range s1 {
		if unicode.Is(unicode.Han, j) {
			sum++
		}
	}
	fmt.Println(sum)
}
