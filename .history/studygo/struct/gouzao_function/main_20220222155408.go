package main

import "fmt"

type game struct {
	name         string
	onlineplayer int
	comment      string
}

func newGame(name string, onlineplayer int, comment string) game {
	return game{
		name:         name,
		onlineplayer: onlineplayer,
		comment:      comment,
	}
}

func main() {
	game1 := newGame("CS GO", 3500000, "NEWBEE")
	fmt.Println("game1:", game1)
}