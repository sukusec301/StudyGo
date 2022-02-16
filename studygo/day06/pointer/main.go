package main

import "fmt"

func main() {
	// a := 10
	// a_pointer := &a
	// a_pointerValue := *a_pointer
	// fmt.Println(a, a_pointer, a_pointerValue)
	var a1 *int
	var a2 = new(int) //分配一块内存并定义一个指针
	//*a = 100
	fmt.Println(a1)
	fmt.Println(a2)

	//new函数申请一个内存地址
}
