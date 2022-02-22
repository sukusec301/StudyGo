package main

import "fmt"

type game struct {
	name    string
	age     int
	comment string
}

func newGame(name string, age int, comment string) game {
	return game{
		name:    name,
		age:     age,
		comment: comment,
	}

}

func main() {
	game1 := newGame("CS GO", 3500000, "zhenbuchuo!")
	fmt.Println("game1 is :", game1)
}
