package main

import "fmt"

func main() {
	m1 := make([]map[int]string, 10, 10)
	m1[0] = make(map[int]string, 1)
	m1[0][10] = "beijing"
	fmt.Println(m1)

	m2 := make(map[string][]string, 10)
	m2["beijing"] = []string{"朝阳区", "海淀区", "东城区", "房山区", "昌平区"}
	m2["shanghai"] = []string{"上海外滩", "莆田区", "老八区"}
	fmt.Println(m2)
}
