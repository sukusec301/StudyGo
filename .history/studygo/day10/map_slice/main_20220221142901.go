package main

import "fmt"

func main() {
	m1 := make([]map[int]string, 0, 10)
	m1[0][100] = "beijing"
	fmt.Println(m1)
}
