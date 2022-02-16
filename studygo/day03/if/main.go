package main

import "fmt"

//if 条件判断
func main() {
	age := 19
	if age > 18 {
		fmt.Println("澳门首家线上赌场开业啦！")
	} else {
		fmt.Println("回去上学吧！")
	}
	//if 多个判断条件
	if age > 35 {
		fmt.Println("人到中年了！")
	} else if age > 18 {
		fmt.Println("青年！")
	} else {
		fmt.Println("好好学习！")
	}

	//if 判断的特殊写法
	if age2 := 30; age2 > 18 {
		fmt.Println("你可真老啊！")
	} else {
		fmt.Println("正直青春，注意身体啊！")
	}
	//fmt.Println(age2)打印无效，是因为age2此时的
	//作用域在if判断里面而不是全局
}
