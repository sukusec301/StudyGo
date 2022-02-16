package main

import "fmt"

func main() {
	i := 1

	for ; i < 10; i++ {
		u := 1
		for ; u <= i; u++ {
			fmt.Printf("%dx%d=%d\t", u, i, i*u)
		}
		fmt.Println()

		// for u := 1; u < 10; u++ {
		// 	fmt.Printf("%d\n", i*u)
		// }
	}
}
