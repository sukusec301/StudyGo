package main //	1、package是进行包的声明，建议声明的包与所在文件夹同名
//main包是程序的入口包，一般main函数会放在这个包下
import (
	"fmt"
	"study/studygo/day07/package/demo01/project_test/dbutils"
)

func main() {
	fmt.Println("这是main包中的main函数执行") //println()也是可以的
	dbutils.DbConnect()              //想要调用函数，必须导入包，并定位到之前所在的包
}
