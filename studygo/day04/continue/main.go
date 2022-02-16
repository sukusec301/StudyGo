// package main

// import "fmt"

// func main() {
// 	for i := 0; i < 10; i++ {
// 		if i == 5 {
// 			continue
// 		}
// 		fmt.Println(i)
// 	}
// 	// fmt.Println()
// }

package main

import "fmt"

/*continue与for循环的变种写法*/
func main() {
	i := 0
	for i < 10 {

		if i == 5 {
			continue
		}
		fmt.Println(i)
		i++
	}
	fmt.Println("Over!")
}
