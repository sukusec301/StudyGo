package main

import (
	"strings"
)

func main() {
	s1 := "How do you do"
	//先遍历出来每个单词，注意，这次不是字符了，而是遍历单词，所以不能用for range遍历
	//所以先切割成切片
	s2 := strings.Split(s1, " ")
	//随后遍历切片，将遍历出来的结果存储到map里面
	m1 := make(map[string]int)
	for _, j := range s2 {
		//1.如果原来map中不存在j这个key，那么出现次数=1
		//2.如果map中存在j这个key，那么出现次数+1
		_, ok := m1[j]
		if !ok {
			m1[j] = 1
		} else {
			m1[j]++
		}
	}

	//再判断是不是第一次出现
	//如果是第一次出现就直接=1
	//如果不是第一次出现，+1

}
