package main

type game struct {
	name         string
	onlineplayer int
	comment      map[string]string
}

func newGame(name string, onlineplayer int, comment map[string]string) game {
	return game{
		name:         name,
		onlineplayer: onlineplayer,
		comment:      comment,
	}
}
