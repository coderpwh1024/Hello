package main

import "fmt"

func main() {

	var a int = 21

	var b int = 20

	if (a == b) {
		fmt.Printf("第一行 - a等于b\n")
	} else {
		fmt.Printf("第一行 - a不等于b\n")
	}

	if (a < b) {
		fmt.Printf("第二行 - a小于b\n")
	} else {
		fmt.Printf("第二行 - a不小于b\n")
	}

	if (a > b) {
		fmt.Printf("第三行 - a大于b\n")
	} else {
		fmt.Printf("第三行 - a不大于b\n")
	}

}