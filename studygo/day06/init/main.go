package main

import (
	"Study_GO/studygo/day06/init/testutils"
	"fmt"
)

var x int = test()

func test() int {
	fmt.Println("test函数被调用！")
	return 10
}

func main() {
	fmt.Println("main will be the 1st!")
	fmt.Println("Name=", testutils.Name, "Gender=", testutils.Gender, "Age=", testutils.Age)
}

func init() {
	fmt.Println("main中的init被执行了")
}
