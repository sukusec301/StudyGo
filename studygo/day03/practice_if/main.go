package main

import "fmt"

func main() {
	age := 35
	// fmt.Scanf("%d",&age)
	if age > 30 {
		fmt.Println("澳门")
	} else if age > 45 {
		fmt.Println("大哥请！")
	} else {
		fmt.Println("未成年，写作业去吧！")
	}
	// fmt.Println()
}
