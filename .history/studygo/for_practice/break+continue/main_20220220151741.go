package main

import "fmt"

func main() {
	var flag == false
	for i := 1; i < 27; i++ {
		for u := 'A'; u <= 'Z'; u++ {
			fmt.Printf("%d %c\n", i, u)
			if i == 5 {
				break
			}
		}
	}
	fmt.Println("It's over!")
}
