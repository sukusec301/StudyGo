package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num1 int8 = 127
	var num2 int32 = 293981269
	var num3 int16 = -23120
	fmt.Printf("%d\n", num1)
	fmt.Printf("%d\n", num2)
	fmt.Printf("%d\n", num3)
	fmt.Println(unsafe.Sizeof(num2)) //查看占几个字节，使用到的是unsafe.Sizeof，需要导入unsafe包

}
