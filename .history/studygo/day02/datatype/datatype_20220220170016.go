package main

import "fmt"

var i1 = 101 //十进制数
//HACK
func main() { //BUG
	fmt.Printf("%d\n", i1) //%d代表输出十进制数
	i2 := 077              //八进制数
	fmt.Printf("%d\n", i2)
	i3 := 0x01234567 //十六进制数
	fmt.Printf("%d\n", i3)
	fmt.Printf("%o\n", i1) //将i1十进制数输出成八进制，八进制设置权限用的多
	fmt.Printf("%x\n", i1) //将i1转换成十六进制数，内存地址的多一些
	fmt.Printf("%b\n", i1) //将i1转换成二进制数
	fmt.Printf("%T\n", i3) //%T表示查看对应的数据类型
	//声明int8类型的变量，int16同理，否则默认为int类型
	i4 := int8(9)
	fmt.Printf("%T\n", i4)
	//以上通通为整型int类型
}
