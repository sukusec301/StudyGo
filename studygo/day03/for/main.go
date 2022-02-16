package main

import "fmt"

func main() {
	//for 循环基本格式，最常用
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// //for 循环变种1：省略初始语句
	// var i = 5
	// for ; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// //for 变种2：省略结束语句，类似于其他语言中的while循环
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }
	// //死循环
	// for {
	// 	fmt.Println("Helloworld!")
	// }

	//for range 循环
	s := "Hello你好"
	for u, v := range s { //u表示索引，v表示字符值
		fmt.Printf("%d %c\n", u, v)
	}
}
