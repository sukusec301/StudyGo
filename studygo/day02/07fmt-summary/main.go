package main

import "fmt"

//fmt 占位符总结

func main() {
	var n = 100
	var s = "我爱你中国"
	//查看类型
	fmt.Printf("%T\n", n)
	fmt.Printf("%#v\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%#b\n", n)
	fmt.Printf("%#o\n", n)
	fmt.Printf("%#x\n", n)
	fmt.Printf("%s\n", s)

}
