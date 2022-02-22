package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Score int
	Hobby []string
}

func main() {
	var a Person
	a.Name = "baizhantang"
	fmt.Println(a.Name)
}
