package main

import "fmt"

func main() {
	var a = [8]int{1, 2, 3, 4, 5, 6, 78, 1}
	slice := a[1:4]
	fmt.Printf("len(slice):%d, 数组容量为:%d", len(slice), cap(slice))

}
