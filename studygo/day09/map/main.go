package main

import "fmt"

func main() {
	//定义map变量：
	var a map[int]string
	//只声明map，内存是没有分配空间的
	//必须通过make函数进行初始化才能分配空间
	a = make(map[int]string, 10) //map可以存放10个键值对
	a[2000] = "张三"
	a[2001] = "李四"
	a[2002] = "王五"
	a[2003] = "赵六"
	a[2001] = "孙七"
	fmt.Println(a)

	//map的声明方式2
	b := make(map[int]string)
	b[20095452] = "张三"
	b[2001] = "孙七"
	fmt.Println("b的长度为：b的值为:", len(b), b)

	//map的声明方式3
	c := map[int]string{
		1: "rnmb",
		2: "cap",
		3: "riledou",
	}
	fmt.Println(c)

	d := map[int]string{
		1: "锄禾日当午",
		2: "汗滴禾下土",
		3: "谁之盘中餐",
		4: "粒粒皆辛苦",
	}
	fmt.Println("更改前的map：", d)
	d[1] = "鹅鹅鹅，曲项向天歌"
	fmt.Println("更改后的map：", d)
}
