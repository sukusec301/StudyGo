package main

import "fmt"

func main() {
	s1 := make([]int, 5, 10)
	s2 := make([]int, 0, 10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v len(s2)=%d cap(s2)=%d\n", s2, len(s2), cap(s2))

	//切片的赋值
	s3 := []int{1, 3, 5} //定义s3这个切片
	s4 := s3             //s3 和 s4都指向了同一个底层数组
	fmt.Println(s3, s4)
	s3[0] = 1000
	s3[1] = 2000
	fmt.Println(s3, s4)

	//切片的遍历，也是支持index索引遍历和range遍历的
	for i, v := range s3 {
		fmt.Printf("%d %d\n", i, v)
	}
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}
}
