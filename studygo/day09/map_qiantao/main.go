package main

import "fmt"

func main() {
	a := make(map[string]map[int]string)

	a["班级1"] = make(map[int]string, 3)
	a["班级1"][201502031] = "beijing"
	a["班级1"][2015132031] = "shanghai"
	a["班级1"][201123123] = "guangdong"
	a["班级1"][201123123] = "shenzhen"

	for k1, v1 := range a {
		fmt.Println(k1)
		for k2, v2 := range v1 {
			fmt.Printf("%v,%v \t", k2, v2)
		}
	}
}
