package main

import "fmt"

func funcA() {
	fmt.Println('a')
}

func funcB() {
	fmt.Println('b')
}

func funcC() {

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
