package dbutils

import "fmt"

func DbConnect() { //注意，如果这是要被调用，需要首字母大写，不然无法导入
	fmt.Println("这是dbutils包下的dbConnect函数执行")
}
