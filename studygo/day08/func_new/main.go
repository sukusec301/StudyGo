package main

import "fmt"

func main() {
	new1 := new(int)
	fmt.Printf("new1的值为：%v, new1的类型为：%T,new1的指针为:%v,new1地址指向的值为：%v", new1, new1, &new1, *new1)
}
