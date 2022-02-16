package main

import (
	"fmt"
	"strings"
)

func main() {
	// \ 本来是具有特殊含义的，我应该告诉程序我写的\就是一个单纯的\
	path := "E:\\精进\\Go语言学习\\01-50"
	fmt.Println(path)
	//fmt.Println("\"E:\\精进\\Go语言学习\\01-50\"")
	s := "I'm ok!"
	fmt.Println(s)

	//多行字符串
	s2 := `			
	世情薄
		人情恶
			雨送黄昏花易落
	`
	fmt.Println(s2)
	s3 := `你好，我是人类高质量男性！`
	fmt.Println(len(s3)) //打印字符串长度，即字节长度，13*3=39

	//字符串拼接
	name := "理想"
	word := "大帅比"
	ss := name + word + "1"
	//用ss1来接收name+word这两个变量
	ss1 := fmt.Sprintf("%s%s+2", name, word)
	fmt.Println(ss)
	fmt.Println(ss1)
	// 分隔，就是path中，去掉\后的结果，类似于一个列表和数组
	ret := strings.Split(path, "\\")
	fmt.Println(ret)
	//ss中是否包含字符串"理想"
	fmt.Println(strings.Contains(ss, "理性"))
	//前缀，即判断是不是理想开头的
	fmt.Println(strings.HasPrefix(ss, "理想"))
	//后缀，即判断是不是理想结尾的
	fmt.Println(strings.HasSuffix(ss, "理想"))
	s4 := "abcdeb"
	//c在字符串出现的位置
	fmt.Println(strings.Index(s4, "c"))
	//a在字符串最后出现的位置
	fmt.Println(strings.LastIndex(s4, "a"))
	//拼接，这里必须是ret，不能是别的变量名
	//这里涉及到了切片，以后再细说哦
	fmt.Println(strings.Join(ret, "+"))
}
