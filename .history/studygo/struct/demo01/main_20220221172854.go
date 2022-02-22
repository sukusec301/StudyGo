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
	a.Hobby = []string{"足球", "篮球", "羽毛球"}
	a.Score = 98
	a.Age = 21
	fmt.Println(a)
	fmt.Println(a.Name)
	fmt.Println(a.Hobby)
	fmt.Println(a.Score)
	fmt.Println(a.Age)
}
