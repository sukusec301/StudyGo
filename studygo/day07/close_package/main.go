package main

import "fmt"

func init() {
	logo := `
	=======================================
	==想念初恋-------------------------=====
	=======================================
	by:sukusec301
	`
	fmt.Println(logo)
}
func getsum() func(int) int {
	var sum int = 0
	return func(num int) int {
		sum = sum + num
		return sum
	}

}

//闭包：返回的匿名函数+匿名函数以外的变量num
func main() {
	f := getsum()
	fmt.Println(f(1))
	fmt.Println(f(1))
	fmt.Println(f(1))
	fmt.Println(f(1))
	fmt.Println("==============我是一条优美的分割线===============")
	// fuck := sum()
	fmt.Println(sum(0, 1))
	fmt.Println(sum(1, 1))
	fmt.Println(sum(2, 1))
}

//不使用闭包来实现一个累加的效果可以吗？
func sum(shu1, shu2 int) int {
	shu1 = shu1 + shu2
	return shu1
}
