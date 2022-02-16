package main

import "fmt"

//找出和为8的两个元素的代表
//定义两个for循环，外层的从第一个开始遍历
//内层的for循环从外层后面的哪个开始找
//他们两个数的和为8
func main() {
	a := [...]int{0, 1, 3, 5, 7, 8}
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}

}
