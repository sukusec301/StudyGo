package main 


import "fmt"


type game struct{
	name string
	age int
	comment mstring
}

func newGame(name string, age int, comment string)game{
	name : name,
	age : age,
	comment : comment,
}

func main(){
	game1 := newGame("CS GO", 3500000,string)
	fmt.Println("game1 is :", game1)
}