/*
 * @Author: sukusec
 * @Date: 2022-02-20 22:16:26
 * @Last Modified by: mikey.zhaopeng
 * @Last Modified time: 2022-02-20 22:22:03
 */
package main

import "fmt"

func main() {
	var m1 map[string]int //这是map初始化
	//fmt.Println(m1 == nil) 返回的是一个true，说明此为空，还未初始化，在内存中没有开辟空间
	m1 = make(map[string]int, 10) //可以自动扩容，但是初始化的时候最好估算好该map的容量，避免在程序运行期间再动态扩容
	//10就是键值对的数量
	m1["beijing"] = 1
	m1["shanghai"] = 2
	m1["jiwuming"] = 35
	fmt.Println(m1)
}
