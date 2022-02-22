package main

import "fmt"

func main() {
	//切片的定义方式1：直接定义
	var slice1 []int
	slice1 = []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)

	//切片的定义方式2：从数组中取出切片
	a1 := [...]int{1, 2, 3, 3, 4, 12, 3}
	s3 := a1[0:4]
	fmt.Println(s3)

	s4 := a1[:4]
	fmt.Println(s4)
	//HACK:这里有些不确定
	//TODO:下次打开后定位的位置
	//NOTE:添加一些说明
	//INFO:Hi
	//TAG:标记一下
	//XXX:随便写写啦
	//BUG:这里有毒
	//FIXME:
}
