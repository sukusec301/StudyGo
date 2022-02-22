package main

import (
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Score int
	Hobby []string
}

//为什么要有结构体？一定是之前所学无法表示或者表示很繁琐新内容了，所以要开发一个新的东西，来解决这个痛点
//解决的就是无法定义一个多维度的东西
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

	//匿名结构体
	var s struct {
		name string
		age  int
	}
	s.name = "beijing"
	s.age = 18
	fmt.Printf("type:=%T\n, value:= %v", s.age, s.age)

}
