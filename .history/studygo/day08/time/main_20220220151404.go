package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("now函数的数值为：%v, now函数输出的类型为%T\n", now, now)

	fmt.Printf("年：%v\n月：%v\n日：%v\n时：%v\n分：%v\n秒：%v\n", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("-------------------------------------------------------------------------------------------------------------------------\n")

	fmt.Printf("当前时间为：%d年%d月%d日 %d时%d分%d秒", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println()
	//拿到函数运行计算并且输出出来的结果，赋予给一个变量，之后我们就可以一直调用这个了。
	date := fmt.Sprintf("当前时间为：%d年%d月%d日 %d时%d分%d秒", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(date)

	date2 := now.Format("2006/01/02 15/04/05")
	date3 := now.Format("2006 15")
	fmt.Println(date2)
	fmt.Println(date3)
}
