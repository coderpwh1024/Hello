package main

import "fmt"

var a2 int = 20

func main() {

	var a2 int = 10

	var b int = 20

	var c int = 0

	fmt.Printf("main() 函数中 a=%d\n", a2)

	c = sum(a2, b)

	fmt.Printf("main() 函数中 c=%d\n", c)
}

func sum(a, b int) int {

	fmt.Printf("sum() 函数中 a=%d\n", a)
	fmt.Printf("sum() 函数中 b=%d\n", b)

	return a + b

}
