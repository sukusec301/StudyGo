package testutils

import "fmt"

var Name string
var Age int
var Gender string

func init() {
	fmt.Println("test中的init函数被执行")
	Name = "你好"
	Age = 18
	Gender = "boy"
}
