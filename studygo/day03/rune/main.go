package main

import "fmt"

func main() {
	//"Hello" => 'H' 'e' 'l' 'l' 'o'
	//字符串修改
	s2 := "白萝卜"
	s3 := []rune(s2) //把字符串强制转换成了一个rune切片
	s3[0] = '红'
	fmt.Println(string(s3))
	//首先go字符串中包含了多个字符，字符可能是一个字节的英文字符byte
	//也可能是像中文这样的三个字节字符rune类型
	c1 := "我"
	c2 := '我'
	fmt.Printf("c1:%T c2:%T", c1, c2)
	c3 := "H"
	c4 := byte('H')
	fmt.Printf("c3:%T c4；%T\n", c3, c4)
	fmt.Printf("%d\n", c4)
}
