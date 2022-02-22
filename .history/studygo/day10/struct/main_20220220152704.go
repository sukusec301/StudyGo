package main

import "fmt"

type test struct {
	TestValue  string
	FuckShit   int
	YourChoice bool
}
type guns struct {
	//变量名大写的话，外界是可以访问到的
	Name    string
	Bullets int
	Danjia  int
}

func main() {
	var g1 guns
	g1.Name = "AK47"
	g1.Bullets = 30
	g1.Danjia = 4
	fmt.Println(g1.Name)
	fmt.Println(g1.Bullets + 100)
	fmt.Println(g1.Danjia + 10)

	var g2 guns
	g2.Name = "AWP"
	g2.Bullets = 10
	g2.Danjia = 3
	fmt.Println(g2.Name)
	fmt.Println(g2.Bullets)
	fmt.Println(g2.Danjia)

	var holyshit test
	fmt.Printf("里面的数据类型为:%T %T", holyshit, g1)
}
