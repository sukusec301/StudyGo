package main

import "fmt"

func main() {
	m1 := make([]map[int]string, 10, 10)
	m1[0] = make(map[int]string, 1)
	m1[0][10] = "beijing"
	fmt.Println(m1)
}
