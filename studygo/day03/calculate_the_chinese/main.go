package main

import "fmt"

func main() {
	s := "hello沙河小王子"
	//fmt.Println(len(s))
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
}
