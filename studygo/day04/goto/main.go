// package main

// import "fmt"

// func main() {
// 	//跳出多层循环
// 	var flag = false
// 	for i := 0; i < 10; i++ {
// 		for j := 'A'; j < 'Z'; j++ {
// 			if j == 'G' {
// 				flag = true
// 				break
// 			}
// 			fmt.Printf("%v-%c\n", i, j)
// 		}
// 		if flag { //if flag 就等于 if flag == true
// 			break
// 		}
// 	}

// }
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'G' {
				goto XX
			}
			fmt.Printf("%v--%c\n", i, j)
		}
	}
XX:
	fmt.Println("It's over!")
}
