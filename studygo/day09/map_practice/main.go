package main

import "fmt"

//map就是定义多个键值对的一种方式，比数组灵活多变
//第一种定义方式：
func main() {
	// var map[int]string
	// fuckme = make(map[int]string,10)
	//第二种定义方式
	mapA := map[int]string{
		1: "fuckc",
		2: "nihao",
		3: "wocao",
	}

	//第三种定义方式：
	b := make(map[int]string)
	b[1] = "hao"
	b[2] = "emm"
	b[3] = "lalalala"
	b[4] = "gogogogogogo"
	fmt.Println(b)
	fmt.Println(mapA)

	mapA[1] = "allalalalalalalalala"
	fmt.Println(mapA)

	delete(b, 2)
	delete(b, 5)
	fmt.Println(b)

	value, flag := b[4]
	fmt.Println(value, flag)
}
