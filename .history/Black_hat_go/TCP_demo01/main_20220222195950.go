package main

import (
	"fmt"
	"net"
)

func main() {

	for i := 1; i < 1025; i++ {
		address := fmt.Sprintf("scanme.nmap.org:", i)
	}

	_, err := net.Dial("tcp", address)
	if err == nil {
		fmt.Println("Connection Completed!!")
	}
}
