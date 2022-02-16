package main

import "fmt"

func main() {
	//算术运算符
	a := 10
	b := 200
	a++ //++在go中表示是一个语句，不能赋值，如：a = a++这样是错的
	b++
	fmt.Printf("%d\n", b+a)
	fmt.Printf("%d\n", b-a)
	fmt.Printf("%d\n", b*a)
	fmt.Printf("%d\n", b/a)
	fmt.Printf("%d\n", b%a)
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)

	/*关系运算符，返回的永远是true或者false*/
	fmt.Println(a == b) //GO语言是强类型语言，只有相同才能比较
	fmt.Println(a != b) //不等于，没有!==这样的写法
	fmt.Println(a <= b)

	/*逻辑运算符&& || */
	//如果年龄大于18岁并且小于60岁...
	age := 4
	if age > 18 && age <= 60 {
		fmt.Print("正值壮年！")
	} else if age <= 18 {
		fmt.Println("未成年呢?回家写作业去吧~")
	} else {
		fmt.Println("活着真好啊！")
	}
	//如果年龄大于60岁或者小于18岁
	if age <= 18 || age >= 60 {
		fmt.Println("老弱病残！")
	}

	//! 即not取反，原来为真输出假，原来为假输出真
	isMarried := false
	fmt.Println(isMarried)
	fmt.Println(!isMarried)
	//位运算符5的二进制为101 2为10
	fmt.Println(5 & 2)   //000	与运算：两个数的二进制对齐后，都为1的才为1，否则为0
	fmt.Println(5 | 2)   //111	或运算：有一个1就为1
	fmt.Println(5 ^ 2)   //111	异或运算：两位不一样才为1
	fmt.Println(5 << 1)  //1010 => 10
	fmt.Println(1 << 10) // 10000000000 => 1024
	fmt.Println(5 >> 2)  //1 => 1
	//赋值运算符
}
