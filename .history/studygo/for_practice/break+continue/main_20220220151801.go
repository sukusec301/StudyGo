package main

import "fmt"

func main() {
	var flag bool
	for i := 1; i < 27; i++ {
		for u := 'A'; u <= 'Z'; u++ {
			fmt.Printf("%d %c\n", i, u)
			if i == 5 {
				flag == true
				break
			}
		}
		if flag {
			break
		}
	}
	fmt.Println("It's over!")
}
