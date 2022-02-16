package main

import "fmt"

func main() {
	var a1 = [3]int{1, 2, 3}
	for i := 0; i < len(a1); i++ {
		fmt.Println(a1[i])
	}
	fmt.Println("=====================================================================")
	for i, v := range a1 {
		fmt.Printf("%d %d\n", i, v)
	}
	fmt.Println("=====================================================================")

	var a2 = [...]string{"shit", "motherfucker", "holyd", "bro"}
	for i, v := range a2 {
		fmt.Printf("%d %s", i, v)
		fmt.Println()
	}
	fmt.Println("=====================================================================")

	a3 := [...]bool{true, true, false}
	fmt.Println(a3)

}
