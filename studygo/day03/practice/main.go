package main

import "fmt"

func main() {
	a := 10
	b := 1.2
	c := 'N'
	d := "你好啊lqy"
	e := true
	fmt.Printf("type:%T, and the value is: %d\n", a, a)
	fmt.Printf("type:%T, and the value is: %f\n", b, b)
	fmt.Printf("type:%T, and the value is: %v\n", c, c)
	fmt.Printf("type:%T, and the value is: %s\n", d, d)
	fmt.Printf("type:%T, and the value is: %v\n", e, e)
	//fmt.Println(e)
}
