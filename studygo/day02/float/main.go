package main

import (
	"fmt"
)

func main() {
	//math.MaxFloat32	//float32最大值
	f1 := 1.234566
	fmt.Printf("%T\n", f1) //默认go语言中的小数都是float64类型
	f2 := float32(1.234566)
	fmt.Printf("%T\n", f2) //显示声明float32类型
	//f1 = f2  错误        //float32类型的值不能直接赋值给float64类型的变量

}
