package main

import "fmt"

func main() {
	var m1 map[int]string
	m1 = make(map[int]string, 5)
	m1[1] = "beijing"
	m1[2] = "guangdong"
	m1[3] = "shenzhen"
	m1[4] = "shanghai"

	fmt.Println(m1[1])
	fmt.Println(m1[2])
	fmt.Println(m1[3])
	fmt.Println(m1[4])
}
