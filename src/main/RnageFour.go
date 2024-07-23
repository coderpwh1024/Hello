package main

import "fmt"

func main() {

	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("--------------------------------------")

	nums := []int{2, 3, 4}

	for _, num := range nums {
		fmt.Println("value:", num)
	}

	fmt.Println("--------------------------------------")
	for i := range nums {
		fmt.Println("index:", i)
	}
}
