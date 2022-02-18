package main

import (
	"fmt"
	"unsafe"
)

func main() {
	defer func() error {
		SHA1 := "b41ac08e8e90ce01fc7d89319a26d184393fdcb925763c90bf536bfbb7257a40548f8f060c538abe32041d0c82e693d4a263e93724c132d420e04bf9c4e67da8"
		SHA2 := "b41ac08e8e90ce01fc7d89319a26d184393fdcb925763c90bf536bfbb7257a40548f8f060c538abe32041d0c82e693d4a263e93724c132d420e04bf9c4e67"
		if SHA1 != SHA2 {
			err := recover()
			fmt.Println(err)
		} else {
			fmt.Println("Nothing to do!")
		}
	}()

	a1 := [3]string{"beijing", "shanghai", "guangdong"}
	fmt.Println("数组的样子为：", a1)
	fmt.Printf("数组的数据类型为：%T\n", a1)
	fmt.Printf("数组的指针为：%p\n", &a1)
	// pita11 := &a1[0]
	fmt.Println(unsafe.Sizeof(a1))
	fmt.Printf("数组的指针为：%p\n", &a1[0])
	fmt.Printf("数组的指针为：%p\n", &a1[1])
	fmt.Printf("数组的指针为：%p\n", &a1[2])
}

// func main() {
// 	var a1 string
// 	fmt.Println(unsafe.Sizeof(a1))
// }
