package main

import "fmt"

type Game struct {
	name, size string
	players    int64
	guns       []string
}

func main() {
	var a Game
	a.name = "CS GO"
	a.players = 3200000
	a.guns = []string{"AK47", "M4A4", "AWP", "SSG", "P90"}

	var b Game
	b.name = "DOTA"
	a.players = 3300000
	a.guns = []string{"COCO", "PA", "SF", "Viper", "TB", "JUGG", "SA", "DK"}

	fmt.Println("DOTA中的英雄都有：\n", a.guns)
	var c struct {
		name   string
		heros  string
		number int
		keys   []string
	}
	c.name = "Popkart"
	c.heros = "黑骑士z7"
	c.number = 50
	c.keys = []string{"E", "Y", "Ctrl+C", "Ctrl+V"}
	fmt.Println(c.keys)
}
