package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	for u := 1; u <= i; u++ {
	// 		fmt.Printf("%d*%d=%d\t", i, u, i*u)
	// 	}
	// 	fmt.Println()
	// }
	s := "123你好啊！"
	// for i, v := range s {
	for _, v := range s { //只打印值，不打印索引，将索引赋值给哑元变量即可
		fmt.Printf(" %c\n", v)
	}
}
