package main

import "fmt"

//递归函数
//计算n=5的阶乘
func f1(n uint64) uint64 {
	//return 5*4!
	return n * f(n-1)
}

func main() {
	ret := f(5)
	fmt.Println(ret)
}
