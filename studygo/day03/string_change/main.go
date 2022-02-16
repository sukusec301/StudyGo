package main

import "fmt"

func main() {

	s := "123456"
	//n := len(s)
	//fmt.Println(n)

	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(s[i])
	// 	fmt.Printf("%c\n", s[i]) //%c:字符
	// }
	for _, c := range s { //从字符串中拿出具体的字符
		fmt.Printf("%c\n", c)
	}

}
