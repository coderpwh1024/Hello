package main

import "fmt"

func main() {

	var ptr *int

	fmt.Printf("ptr 的值为: %x\n", ptr)

	if ptr != nil {
		fmt.Println("当前不是空指针")
	}

	if ptr == nil {
		fmt.Println("当前是空指针")
	}
}
