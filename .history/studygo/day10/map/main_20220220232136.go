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
	//TODO 继续练习map
	m2 := map[int]string{
		1: "beijing",
		2: "shanghai",
		3: "guangdong",
	}
	fmt.Printf("m2的类型为:%T,m2的值为:%v\n", m2, m2)
	delete(m2, 1)
	delete(m2, 5)
	for key, value := range m2 {
		fmt.Printf("键为：%d,值为：%v\n", key, value)
	}
	fmt.Println(m2)
}
