package main

import "fmt"

func main() {
	var flag bool
	for i := 1; i < 27; i++ {
		for u := 'A'; u <= 'Z'; u++ {
			fmt.Printf("%d %c\n", i, u)
			if i == 5 {
				flag = true
				break
			}
		}
		if flag {
			break //跳出外层的for循环
		}
	}
	fmt.Println("It's over!")
}
