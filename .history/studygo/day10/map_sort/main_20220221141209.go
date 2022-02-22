package main

import (
	"fmt"
	"math/rand"
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
	fmt.Println(scoreMap)

}
