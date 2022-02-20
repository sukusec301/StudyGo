package main

import "fmt"

func init() {
	logo := `
======================================================	
	-=-=-=-=-=-=-									 =
	想念初恋——御剑终端								   =
	-=-=-=-=-=-=-=									 =
======================================================
	`
	fmt.Println(logo)
}
func main() {
	var suibian string
	fmt.Scan(&suibian)
	fmt.Println(suibian, 10)
}
