package main

import "fmt"

type dog struct {
	name string
	age  int
	fuck []string
}

func newDog(name string, age int, fuck []string) dog {
	return dog{
		name: name,
		age:  age,
		fuck: fuck,
	}
}
func main() {
	d1 := newDog("fuckingbitch", 18, []string{"2", "1", "1", "123"})
	fmt.Println(d1)
}
