package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Score int
	Hobby []string
}

func main() {
	Person.Name = "baizhantang"
	fmt.Println(Person.Name)
}
