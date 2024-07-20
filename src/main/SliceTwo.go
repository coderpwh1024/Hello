package main

import "fmt"

func main() {

	var numbers []int

	if numbers == nil {
		fmt.Printf("切片是空的")
	}

}

func printSliceTwo(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)

}
