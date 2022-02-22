package main

import "fmt"

type dog struct {
	name string
}

//构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

//这个函数谁都可以调用，所以叫做函数
//如果只能作用于某种特定的类型才能调用，这时候就叫做方法
//接收者默认使用类型首字母小写来表示，如dog的d。python用多的人可能会写成self
//php开发或者其他语言开发的可能会写成this

func (d dog) wang() { //(d dog)就是接收者
	fmt.Printf("%s:wangwangwang~", d.name)
}
func main() {
	d1 := newDog("fuckingbitch")
	d1.wang() //调用方法

}
