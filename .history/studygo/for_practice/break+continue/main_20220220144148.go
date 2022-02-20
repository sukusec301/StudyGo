package main

import (
	"fmt"
)

func main() {
	// var a int = 5
	// switch a s{
	// case 1:
	// 	fmt.Println("大拇指")
	// case 2:
	// 	fmt.Println("食指")
	// case 3:
	// 	fmt.Println("中指")
	// case 4:
	// 	fmt.Println("无名指")
	// case 5:
	// 	fmt.Println("小拇指")
	// default:
	// 	fmt.Println("不存在你妈的")

	// }//
	var b int
	fmt.Scan(&b)
	switch b {
	case 1, 3, 5, 7, 9:
		fmt.Println("您输入的是奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("您输入的是偶数")
	default:
		fmt.Println("您输入的数超过了10，不好意思请重新输入！")
	}
}
