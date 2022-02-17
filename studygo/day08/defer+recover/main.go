package main

import "fmt"

func main() {
	test()
	fmt.Println("上面的除法操作执行成功。。。")
	fmt.Println("正常执行下面的逻辑！")

}
func test() {
	//利用defer+recover来捕获错误：defer后加上你匿名函数的调用
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("错误已经捕获")
			fmt.Println("err是：", err)
		}
	}()
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)
}
