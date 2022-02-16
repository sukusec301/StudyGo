package main

import "fmt"

func main() {
	fmt.Println("你好，世界\bbb de ") //退一格
	fmt.Println("a\tbb de")      //输出一个水平制表符，8个字符为一个制表符，如果前面有3个字符，空5个格
	fmt.Println("aaaa\tbb de")
	fmt.Println("你好，世界\r bb de") //将光标回到本行开头，后续输入就会替换原有字符
	fmt.Println("你好，世界\"bb de")
	fmt.Println("你好，世界\fbb de")
	fmt.Println(`'`)
}
