package main

import "fmt"

func main() {
	for i := 1; i < 27; i++ {
		for u := 'A'; u <= 'Z'; u++ {
			fmt.Printf("%d %c", i, u)
		}
	}
	fmt.Println("It's over!")
}
