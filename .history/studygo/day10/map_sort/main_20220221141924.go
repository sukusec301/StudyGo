package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	scoreMap := make(map[string]int)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("Stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	// fmt.Println(scoreMap)
	//将结果打印到切片中
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//随后对切片进行排序
	sort.Strings(keys)
	//按照排序后的keys进行遍历

}
