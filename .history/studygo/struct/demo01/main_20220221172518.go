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
	Person.Name = "baizhantang"
	fmt.Println(Person.Name)
}
