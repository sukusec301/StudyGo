package main

import (
	"fmt"
	"strconv"
	"strings"
)

func cal() {
	str := "hello你好！"
	//方法一：for循环遍历字符串
	for i := 0; i < len(str); i++ {
		fmt.Printf("字符串的键值为：%c\n", str[i])
	} //这样的for循环不可能打印出字符串中的值。

	fmt.Println("-----------------------------------------------------")

	//方法二：利用for-range遍历字符串
	for i, j := range str {
		fmt.Printf("字符串的键为：%d,字符串的值为%c\n", i, j)
	}
	fmt.Println("-----------------------------------------------------")
	//利用r := []rune(str)
	//方法三：将字符串转换成切片格式
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c\n", r[i]) //%c就是一个unicode字符
	}

	fmt.Println(len(str)) //len()得出的结果为字节，一个英文字母符号
	//为一个字节，汉字在utf-8中是三个字节

	//字符串转整数
	str1 := strconv.Itoa(88)
	fmt.Println(str1)

	//整数转字符串
	num1, _ := strconv.Atoi("66")
	fmt.Println(num1)

	fmt.Println("----------------------------------------------------")
	fmt.Println("----------------------------------------------------")
	//查看字符串中有多少子串
	str2 := "I can can a can by a can."
	count := strings.Count(str2, "can")
	fmt.Println(count)
	fmt.Println("----------------------------------------------------")
	//不区分大小写进行比较，适用场景就是比较密码或者返回包等数据
	ef := strings.EqualFold("dje8yHDEUy8JIS./2123@##451D", "dje8yHDEUy8IS./2123@##451D")
	fmt.Println(ef) //false，输出一个bool值
	fmt.Println("---------------")

	//返回子串在字符串中第一次出现的索引数
	fmt.Println(strings.Index("渗透测试工程师NEWX666", "6"))

	//返回子串在字符串中最后一次出现的索引数
	fmt.Println(strings.LastIndex("渗透测试工程师NEWX666", "6"))

	//字符串替换
	fmt.Println(strings.Replace("渗透测试渗透工程渗透师NEWX666", "渗透", "肾透", 2))

	//按照某个字符切割

	fmt.Println(strings.Split("求求-哥哥-带带-弟弟", "-"))
	//fmt.Println(strings.Replace("求求-哥哥-带带-弟弟", "求求-哥哥-带带-弟弟", "qqggdddd", -1))

	//字符串字母大小写转换
	fmt.Println(strings.ToLower("FUCkYoU"))
	fmt.Println(strings.ToUpper("fuckyou"))

	//将字符串左右两边的空格去除
	fmt.Println(strings.TrimSpace("    dddd    "))

	//将字符串两边指定的字符去除
	fmt.Println(strings.Trim("~root@localhost~", "~"))

	//将字符串右边指定的字符去掉
	fmt.Println(strings.TrimRight("~root@localhost~", "~"))

	//将字符串左边指定的字符去掉
	fmt.Println(strings.TrimLeft("~root@localhost~", "~"))

	//判断字符串是否以指定的字符串开头

	fmt.Println(strings.HasPrefix("https://github.com/sukusec301/StudyGo.git", "https"))

}
func main() {
	cal()
}
