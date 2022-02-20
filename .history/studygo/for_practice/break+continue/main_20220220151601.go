package main

import "fmt"

func main() {
	for i := 1; i < 27; i++ {
		for u := 'A'; u <= 'Z'; u++ {
			fmt.Println(i, u)
		}
	}
}
