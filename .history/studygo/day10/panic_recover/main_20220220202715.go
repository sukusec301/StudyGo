package main

import "fmt"

func funcA() {
	fmt.Println('a')
}

func funcB() {

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
