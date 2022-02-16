package main

import "fmt"

func main() {
	// var a = 5
	// if a == 1 {
	// 	fmt.Println("大拇指")
	// } else if a == 2 {
	// 	fmt.Println("食指")
	// } else if a == 3 {
	// 	fmt.Println("中指")
	// } else if a == 4 {
	// 	fmt.Println("无名指")
	// } else if a == 5 {
	// 	fmt.Println("小拇指")
	// }

	//使用switch来简化上述代码
	a := 6
	switch a {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("食指")
	case 4:
		fmt.Println("食指")
	case 5:
		fmt.Println("食指")
	default:
		fmt.Println("不存在")
	}
}
