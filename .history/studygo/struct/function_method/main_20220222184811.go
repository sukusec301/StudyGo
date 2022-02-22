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

//这个函数谁都可以调用，所以叫做函数
//如果只能是某个特定的类型才能调用，这时候就叫做方法
func wang() {
	fmt.Println("wwwwwww！")
}
func main() {
	d1 := newDog("fuckingbitch")
	fmt.Println(d1)
}
