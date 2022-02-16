package main

import "fmt"

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)
	m1["理想"] = 18
	m1["你好"] = 35

	fmt.Println(m1)
}
