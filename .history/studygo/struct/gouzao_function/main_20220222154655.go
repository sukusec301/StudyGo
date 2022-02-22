package main

import "fmt"

type person struct {
	name string
	age  int
}
type dog struct {
	name string
}
type cat struct {
	name string
	age  int
	fuck []string
}

func newCat(name string, age int, fuck []string) cat {
	return cat{
		name: name,
		age:  age,
		fuck: fuck,
	}
}
func main() {
	cat1 := newCat("shit", 18, []string{"体弱", "多病", "的小猫咪"})
	fmt.Println("cat1是：", cat1)
	// var p person
	// p.name = "nihao"
	// p.age = 18
	//以往我们会进行上面的声明，但是这次不同，我们使用函数
	p1 := newPerson("令狐冲", 30)
	p2 := newPerson("华山派", 300)
	fmt.Println(p1, p2)

	dog1 := newDog("飞子")
	fmt.Println("这条狗的名字叫做：", dog1)
}

func newPerson(name string, age int) person {
	return person{
		name,
		age,
	}
}
func newDog(name string) dog {
	return dog{
		name: name,
	}
}
