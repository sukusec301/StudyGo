package main

import "fmt"

func main() {
	// a := [3][2]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	// }
	// fmt.Println(a)
	// fmt.Println(a[2][1])
	/*二维数组遍历*/

	a := [3][2]string{
		{"beijing", "shanghai"},
		{"yingguo", "lundun"},
		{"huangtudi", "shuaxin"},
	}
	for _, v := range a {
		for _, v1 := range v {
			fmt.Printf(" %s\n", v1)
			fmt.Println("================")
		}

	}

}
