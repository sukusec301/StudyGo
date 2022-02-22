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
	cat1.cao()
}

func (c cat) cao() {
	s1 := "你肯定猜不到我会是一只猫，一直陪在你的身边"
	for _, j := range s1 {
		fmt.Printf("这句话的每个字分别为：\n", j)
	}
}
