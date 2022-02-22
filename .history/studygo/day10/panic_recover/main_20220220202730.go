package main

import "fmt"

func funcA() {
	fmt.Println('a')
}

func funcB() {
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
