package main

import "fmt"

type dog struct {
	name string
}

//构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}
func main() {
	d1 := newDog("fuckingbitch")
	fmt.Println(d1)
}
