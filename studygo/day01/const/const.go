package main

import "fmt"

//声明常量
const pi = 3.1415926

//批量声明常量
const (
	li = 1
	yi = 2
	ni = 3
)

// 假设批量声明时没有定义其他的值，默认和上面的一样
const (
	fi = 100000
	fe
	fee
)

//iota类似于索引值，遇到const，重置为0
const (
	c1 = iota //0
	c2 = iota //1
	c3 = iota //2
	c4        //3 因为我没有写值，默认和上面的一样，即iota，但是累加了一个，即2+1=3
	c5        //4
	c6        //5
)
const (
	b1 = iota //0 是因为iota一旦遇到const，就重置为0
	b2 = iota //1
	_         //2	什么都没写，即_ = iota 1+1=2
	b3        //3   同理，什么都没写，即b3 = iota 2+1=3
)
const (
	_  = iota             //0
	KB = 1 << (10 * iota) //1往前移动了10位，即2的10次方，1024
	MB = 1 << (10 * iota) //
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	a := 1 //复习，简短变量声明法
	fmt.Println(a)
	fmt.Println(fe + fee)
	fmt.Println("hello", ni)
	fmt.Println("hello", yi)
	fmt.Println("nihao", pi)
	fmt.Println("iota")
	fmt.Println("c1", c1)
	fmt.Println("c2", c2)
	fmt.Println("c3", c3)
	fmt.Println("c4", c4)
	fmt.Println("c5", c5)
	fmt.Println("c6", c6)
	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("b3:", b3)
	fmt.Println("KB", KB)
	fmt.Println("MB", MB)
	fmt.Println("GB", GB)
	fmt.Println("TB", TB)
	fmt.Println("PB", PB)
}
