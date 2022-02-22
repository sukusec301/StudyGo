package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "How do you do"
	//先遍历出来每个单词，注意，这次不是字符了，而是遍历单词，所以不能用for range遍历
	fmt.Println(strings.Split(s1, " "))
	//再判断是不是第一次出现
	//如果是第一次出现就直接=1
	//如果不是第一次出现，+1

}
