// package main

// import "fmt"

// func main() {
// 	a1 := []int{1, 3, 5}
// 	a2 := a1
// 	//var a3 []int //这么写相当于切片a3 == nil为空，肯定是将a1拷贝不进去
// 	var a3 = make([]int, 3) //这样子写就正确了，其实make函数不用写后面的容量都是可以的
// 	copy(a3, a1)
// 	fmt.Println(a1, a2, a3)

// 	a1[0] = 2000 //更改a1，那么a2也会发生变化，因为他们两个用的同一块内存
// 	fmt.Println(a1, a2, a3)
// }

package main

import "fmt"

var (
	a1 = []int{1, 2, 3, 4, 5, 2, 6, 7, 8, 8}
)

func main() {

	//删除上面切片的5，注意是删除切片的5！
	a1 := append(a1[:4], a1[5:]...) //第一个当然不需要拆开，第一个是我将要传入的切片
	fmt.Println(a1)
	fmt.Println(cap(a1)) //容量还是为10，因为底层数组就是10个，切片永远不存储值，只是相当于操作数组的一个草图！
	/*如果还是不理解，看看下面这段分析*/
	array := [...]int{1, 2, 3, 4} //定义个数组
	slice := array[:]             //切片的第二种定义方式：从数组定义切片
	fmt.Println(len(slice), cap(slice))
}
