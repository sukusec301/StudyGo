package main

type person struct {
	name string
	age  int
}

func main() {
	// var p person
	// p.name = "nihao"
	// p.age = 18
	//以往我们会进行上面的声明，但是这次不同，我们使用函数

}

func newPerson(name string, age int) person {
	return &person{
		name: name,
		age:  age,
	}
}
