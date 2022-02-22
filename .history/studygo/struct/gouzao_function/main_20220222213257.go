package main

type dog struct {
	name string
	age  int
}

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

func main() {
	newPerson("shiqigege", 24)
}
