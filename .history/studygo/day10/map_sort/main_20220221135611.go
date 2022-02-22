package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ran := rand.Seed(time.Now().UnixNano()) //UnixNano拿到当前时间的纳秒数后6位 然后加一个随机数种子
	var scoreMap = make([string]int)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%0ad", i)
	}
}
