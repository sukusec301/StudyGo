package main

import "fmt"

func main() {
	/*当i==5时，跳出循环*/
	i := 0
	for i < 10 {
		fmt.Println(i)
		if i == 5 {
			break
		}
		i++
	}
	fmt.Println("Over!")

}
