package main

import "fmt"

type cat struct {
	name string
	age  int
}

func newCat(name string, age int) cat {
	return cat{
		name: name,
		age:  age,
	}
}

func (c cat) miao() {
	fmt.Printf("%s\n", c.name)
}
func main() {
	cat1 := newCat("豆", 5)
	fmt.Println(cat1)
	fmt.Printf(cat1.miao())

}
