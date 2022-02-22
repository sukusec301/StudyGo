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
	fmt.Printf("这只猫的名字为：%s\n", c.name)
	fmt.Println("我们一起学猫叫，一起喵喵喵喵喵！")
}
func main() {
	cat1 := newCat("豆", 5)
	fmt.Println(cat1)
	cat1.miao()
}
