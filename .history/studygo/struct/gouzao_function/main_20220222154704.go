package main 


import "fmt"


type game struct{
	name string
	onlinePlayer int
	comment map[string]string
}

func newGame(name string, onlinePlayer int, comment map[string]string)game{
	name : name,
	onlinePlayer : onlinePlayer,
	comment : comment,
}

func main(){
	game1 := newGame("CS GO", 3500000,map[string]string)
	fmt.Println("game1 is :", game1)
}