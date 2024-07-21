package main

import "fmt"

func main() {

	var numbers []int

	printSliceFour(numbers)

	numbers = append(numbers, 1)
	printSliceFour(numbers)

	numbers = append(numbers, 2, 3, 4)
	printSliceFour(numbers)

	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	copy(numbers1, numbers)
	printSliceFour(numbers1)

}

func printSliceFour(x []int) {

	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)

}
