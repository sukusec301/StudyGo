package main

import "fmt"

func funcA() {
	fmt.Println('a')
}

func funcB() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("错误！错误正在回收中，请稍等！")
		}
	}()
	panic("fuckyou!")
	fmt.Println('b')
}

func funcC() {
	fmt.Println('c')
}

func funcD() {
	fmt.Println('D')
}

func main() {
	funcC()
	funcA()
	funcB()
	funcD()
}
