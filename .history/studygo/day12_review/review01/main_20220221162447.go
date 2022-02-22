package main

import "fmt"

func yuanshuai(name string) {
	fmt.Println("Hellow!", name)
}

func lixiang(f func(string)) {
	f()
}

func main() {
	lixiang(yuanshuai)
}
