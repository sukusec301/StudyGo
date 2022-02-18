package main

import (
	"fmt"
)

func main() {
	//数组，即存放元素的容器
	//必须指定存放的元素的类型和容量（长度）
	//注意，数组的长度是类型的一部分，所以下方这两组数组无法进行比较，就是因为两者为不同的类型。
	var a1 [3]bool //[true false true]
	var a2 [4]bool //[true false true false]
	fmt.Printf("a1:%T \na2:%T", a1, a2)

	//数组的初始化
	//如果不初始化，默认元素都是零值（零值就是各种数据类型的默认数值。即bool=false，string=""，int=0，float=0）
	fmt.Println(a1, a2)
	//数组初始化第一种方式：
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	//数组初始化第二种方式：未知数量初始化，根据推断判断数组的长度多少
	a3 := [...]int{1, 1, 1, 4, 3, 12, 32, 3, 23, 23, 214, 1, 41, 5, 1}
	fmt.Println(a3)
	//方式3：根据索引来初始化，第一位为1，最后一位是2
	a4 := [5]int{0: 1, 4: 2}
	fmt.Println(a4)
	//特殊变种，长度为10的数组，里面只写2个，剩下的都会被零值补全
	a5 := [10]int{1, 2}
	fmt.Println(a5)
	test()
}

func test() {
	//数组：就是存放元素的容器，必须指定存放原色的类型和容量（长度）
	var a1 = [...]string{"dea", "dead", "dasd", "asda"}
	fmt.Println(a1[0], a1[1], a1[2], a1[3])
	var a2 = [2]int{1, 2}
	fmt.Println(len(a2), cap(a2))
	var a3 = [...]int{213, 21, 312, 3, 213, 123, 12, 3}
	for i := 0; i < len(a3); i++ {
		sum := 0
		sum += a3[i]
	}
	a4 := [...]string{"10", "123", "213", "123", "123", "89673"}
	sum := 0
	for _, j := range a4 {
		sum = sum + j
		fmt.Printf("数组里面数的和为：%d", sum)
	}

}
