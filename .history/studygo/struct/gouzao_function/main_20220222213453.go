package main

import "fmt"

// type dog struct {
// 	name string
// 	age  int
// }

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}
func sf(name string, age int) {
	p1.age++
}
func main() {
	p1 := newPerson("shiqigege", 24)
	fmt.Println(p1.age)

	fmt.Println(p1.age)
}
