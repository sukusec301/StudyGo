package main

import "fmt"

func main() {
	// var m1 map[int]string
	// m1 = make(map[int]string, 5)
	// m1[1] = "beijing"
	// m1[2] = "guangdong"
	// m1[3] = "shenzhen"
	// m1[4] = "shanghai"

	// fmt.Println(m1[1])
	// fmt.Println(m1[2])
	// fmt.Println(m1[3])
	// fmt.Println(m1[4])

	m2 := make(map[int]string, 8)
	m2[1] = "佟湘玉"
	m2[2] = "吕秀才"
	m2[3] = "郭芙蓉"
	m2[4] = "白展堂"
	m2[5] = "李大嘴"
	fmt.Printf("m1的类型为：%T，m1的值为：", m2, m2)
}
