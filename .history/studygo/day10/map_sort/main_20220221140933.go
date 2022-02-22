package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	m1 := make(map[string]int,10)
	for key,value := range m1{
		m1[key] := fmt.Sprintf("Stu%02d",key) 
	}
}
