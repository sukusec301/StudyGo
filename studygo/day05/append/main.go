package main

import "fmt"

func main() {
	/*append：给切片添加元素*/
	s1 := []string{"beijing", "shanghai", "guangzhou"}
	// s1[3] = "shenzhen" //错误！索引越界！
	// fmt.Println(s1)
	fmt.Printf("len(s1): %d, cap(s1): %d", len(s1), cap(s1))

	fmt.Printf("\n")

	s1 = append(s1, "shenzhen", "hangzhou")
	fmt.Printf("len(s1):%d, cap(s1): %d", len(s1), cap(s1))

	fmt.Printf("\n")

	ss := []string{"chengdu", "chongqing", "suzhou"}
	s1 = append(s1, ss...)
	fmt.Printf("len(s1):%d, cap(s1): %d", len(s1), cap(s1))
}
