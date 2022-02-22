package main

import "fmt"

type dog struct {
	name   string
	age    int
	method string
}
type person struct {
	name   string
	age    int
	method string
}

type game struct {
	name   string
	number int
	player int
}

func newDog(name string, age int, method string) dog {
	return dog{
		name:   name,
		age:    age,
		method: method,
	}
}
func newPerson(name string, age int, method string) person {
	return person{
		name:   name,
		age:    age,
		method: method,
	}
}
func newGame(name string, number int, player int) game {
	return game{
		name:   name,
		number: number,
		player: player,
	}
}
func (d dog) wang() {
	fmt.Println("Dog's method is wangwangwang!", d.name)
	fmt.Println("Dog's method is wangwangwang!", d.age)
}
func (p person) haha() {
	fmt.Println("Person's method is hahahahah!", p.name)
}
func main() {
	d1 := newDog("dogname", 18, "wangwang")
	fmt.Println("d1的值为：", d1)
	p1 := newPerson("personname", 19, "holyshit")
	fmt.Println("p1的值为:", p1)
	d1.wang()
	// p1.wang()
	p1.haha()

}
