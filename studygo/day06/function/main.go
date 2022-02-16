/*
 * @Author: mikey.zhaopeng
 * @Date: 2022-02-13 20:55:45
 * @Last Modified by: mikey.zhaopeng
 * @Last Modified time: 2022-02-13 21:05:00
 */
package main

import "fmt"

//函数
//函数存在的意义？是一段代码的封装
//把一段逻辑抽象出来封装到一个函数中，给它起个名字
//每次用到它的时候，直接调用函数名
//使用函数能够让代码结构更清晰、简洁

//函数的定义
//求两个int之和
func sum(x int, y int) (hello int) {
	return x + y

}

//main函数默认是没有返回值的，所以只有一个参数的括号
func main() {

	hello := sum(101, 999)
	fmt.Println(hello)
}
