package main //开始要有一个包的声明

import "fmt"

//函数外只能放置标识符（变量\常量\函数\类型）的声明

//程序的入口函数
func main() {
	var (
		a int    = 10
		b int    = 1
		c string = "caonima"
		d bool   = true
	)
	const (
		iota = 0
		ee   = iota
		bb
	)

	fmt.Println("Helloworld!")
	fmt.Printf("%d", a)
	fmt.Println()
	fmt.Printf("%s", c)
	fmt.Print("\n")
	fmt.Printf("%d", b)
	fmt.Print("\n")
	fmt.Print(d)
	fmt.Println()
	fmt.Print(bb)
}

//compile this codefile! You'll get hello-world!
