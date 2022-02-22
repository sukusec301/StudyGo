package main

import "fmt"

func main() {
	type Map map[string][]int      //定义Map的类型为 map[string][]int，也就是值为切片的map类型
	m := make(Map)                 //初始化一个Map，赋予变量给m
	s := []int{1, 2}               //s定义为一个具有两个值的切片
	s = append(s, 3)               //添加一个数值，即s = []int{1,2,3}
	fmt.Printf("%+v\n", s)         //打印s
	m["q1mi"] = s                  //m的qlmi键对应值为s
	s = append(s[:1], s[2:]...)    //删除s[1]=2
	fmt.Printf("%+v\n", s)         //s={1,3}
	fmt.Printf("%+v\n", m["q1mi"]) //m["qlmi"]=[]int{1,2,3}
}
