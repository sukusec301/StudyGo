package main

import "fmt"

//Go语言中推荐使用小驼峰式命名
// var student_name string	下划线
// var studentName string	小驼峰，第二个单词首字母大写，推荐！
// var StudentName string	大驼峰，单词首字母都大写

//声明变量
// var name string
// var age int
// var isok bool

//批量声明变量，变量声明后值为空
var (
	name string // ""
	age  int    // 0

	isok bool // false

)

func main() {
	name = "理想"
	age = 16
	isok = true
	//Go语言中非全局变量声明必须要使用，不使用就编译不过去，会变黄颜色。是为了减少空间
	fmt.Println(age)              //打印完指定的内容后，在后面加个换行
	fmt.Print(isok)               //在终端中输出打印的内容
	fmt.Println()                 //快捷打印一个空行，因为Println默认后面加个换行
	fmt.Printf("name:%s\n", name) //%s:占位符，使用name这个变量的值去替换占位符，就是找个人先帮我顶着，等我办完事儿就我就过来了
	fmt.Print("Do you love me?")
	//声明变量的同时赋值，不推荐
	//var s1 string = "whoami"
	//变量声明赋值，类型推导，根据赋值推导变量类型
	//var myname = "sukusec301"
	//简短变量声明，等同于var age = 18。必须在函数体里面使用，不能在全局使用
	//age := 18
	//age := 19，同一个作用域{}中不能重复声明同名的变量
	//匿名变量是一个特殊的变量：_后面学了函数在说
}
