package main

import "fmt"

func main() {

	var sum int = 17

	var count int = 5

	var mean float32

	mean = float32(sum) / float32(count)

	fmt.Printf("mean 的值为:%f\n", mean)

	var a int = 10

	var b float64 = float64(a)

	fmt.Println("b的值为:", b)

}
