package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	m1 := make(map[string]int, 10)
	for key, value := range m1 {
		fuck := fmt.Sprintf("Stu%02d", key)
		valueFuck := rand.Int(100)
		m1[fuck] = valueFuck
	}
	fmt.Println(m1)
}
