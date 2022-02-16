package main

import "fmt"

func main() {
	//切片的定义
	var s1 []int    //定义一个存放int类型元素的切片
	var s2 []string //定义一个存放string类型元素的切片
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) //true
	fmt.Println(s2 == nil) //true
	//初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"beijing", "shanghai", "shenzhen"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) //false
	fmt.Println(s2 == nil) //false	这两个切片已经初始化了，所以就不为空了肯定
	fmt.Printf("s1长度为：%d,s1容量为：%d\n", len(s1), cap(s1))
	fmt.Printf("s2长度为：%d,s2容量为：%d\n", len(s2), cap(s2))
	//2.由数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	s3 := a1[0:4] //基于一个数组切割，左包含右不包含（左闭右开）
	s4 := a1[1:6] //从index=1开始，切到index=6
	s5 := a1[:4]  //相当于从0切到4，包含0
	s6 := a1[3:]  //从3，切到最后
	s7 := a1[:]
	fmt.Println("s3为:", s3)
	fmt.Println("s4为：", s4)
	fmt.Println("s5为：", s5)
	fmt.Println("s6为：", s6)
	fmt.Println("s7为：", s7)
	fmt.Printf("len为：%d，cap为：%d\n", len(s5), cap(s5)) //len:4,cap:8  切片的容量，就是底层数组的容量，即a1的容量为8
	fmt.Printf("len为：%d，cap为：%d\n", len(s4), cap(s4)) //len:5,cap:7  按照上面的说法，这里怎么就是7了？不应该是8么？
	//就是因为，切片是指向底层数组的一个概念，容量即：数组从切片开始的第一个数算起，到数组后的所有数值数量就是容量

	/*切片再切割*/
	s8 := s6[1:3]                                                //s6=[4,5,6,7,8]
	fmt.Printf("s8为：%v，len为：%d，cap为：%d\n", s8, len(s8), cap(s8)) //s8是从5开始的，并且从s6算起，而s6的数组为a1，所以s8的容量就是a1从5到最后：4

}
