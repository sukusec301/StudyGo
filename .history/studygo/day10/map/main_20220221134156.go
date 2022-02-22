package main

import "fmt"

func main() {
	// var m1 map[int]string	//BUG 不知道为什么一直飘黄，可能默认就不支持吧
	// m1 = make(map[int]string, 5)
	// m1[1] = "beijing"
	// m1[2] = "guangdong"
	// m1[3] = "shenzhen"
	// m1[4] = "shanghai"

	// fmt.Println(m1[1])
	// fmt.Println(m1[2])
	// fmt.Println(m1[3])
	// fmt.Println(m1[4])
	//TODO 继续练习map
	m2 := map[string]int{
		"fuck": 10,
		"cao":  1,
		"shit": 12,
	}
	fmt.Println(m2)
	m3 := make(map[int]string, 10)
	m3[1] = "C语言"
	m3[2] = "汇编语言"
	m3[3] = "Java"
	m3[4] = "Go"
	m3[5] = "Python"
	m3[6] = "PHP"
	delete(m3, 3)
	fmt.Println(m3)
}
