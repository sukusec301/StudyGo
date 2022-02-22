package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	// var p person

	// fmt.Printf("%T\n", p)  //p的类型为main.person
	// fmt.Printf("%v\n", p)  //p的值为{0 }
	// fmt.Printf("%p\n", &p) //p的指针为0xc000004078
	var p1 = new(person)
	fmt.Printf("%T\n", p1)  //p1的类型为*main.person
	fmt.Printf("%#v\n", p1) //p1的值为&{ 0}
	fmt.Printf("%p\n", p1)  //p1的值为0xc000004078

}
