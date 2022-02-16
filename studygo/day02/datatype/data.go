package main

import "fmt"

func main() {
	var (
		a = 100
	)
	fmt.Printf("此输出为二进制数：%b\n", a)
	fmt.Printf("此输出为八进制数：%o\n", a)
	fmt.Printf("此输出为十进制数：%d\n", a)
	fmt.Printf("此输出为十六进制数：%x\n", a)
}
