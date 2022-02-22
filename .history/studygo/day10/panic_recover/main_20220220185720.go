package main

import "fmt"

func funcA() {

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
