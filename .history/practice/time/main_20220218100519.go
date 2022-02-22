package main

import (
	"errors"
	"fmt"
	"time"
)

func init() {
	logo := `

╔═╗┬ ┬┌┐┌╦  ┌─┐┌─┐┬┌┐┌   ╦═╗┌─┐┌─┐
╚═╗│ ││││║  │ ││ ┬││││───╠╦╝│  ├┤ 
╚═╝└─┘┘└┘╩═╝└─┘└─┘┴┘└┘   ╩╚═└─┘└─┘

						by:T00ls.net
						向日葵Rce
----------------------------------------------
`
	fmt.Println(logo)

}
func test() error {

	num1 := 10

	num2 := 0

	if num2 == 0 {
		err := errors.New("输入的除数不正确！")
		return err
	}
	result := num1 / num2
	fmt.Println(result)
	return nil
}

func main() {
	date := time.Now().Format("2006/01/02 15/04/05")
	fmt.Println("当前时间为：", date)
	ERR := test()
	if ERR != nil {
		fmt.Println("自定义错误为：", ERR)
		panic(ERR)
	}

}
