package main

import (
	"errors"
	"fmt"
)

func main() {
	suibian := test()
	if suibian != nil {
		fmt.Println("自定义错误：", suibian)
		panic(suibian)
	}

	fmt.Println("asd")
}

func test() error {
	num1 := 10
	num2 := 0
	if num2 == 0 {
		err := errors.New("不能为0！")
		return err
	}
	result := num1 / num2
	fmt.Println(result)
	return nil
}
