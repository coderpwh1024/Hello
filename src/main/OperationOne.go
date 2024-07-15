package main

import "fmt"

func main() {

	var a int = 21

	var b int = 10

	var c int

	c = a + b

	fmt.Printf("第一行 -c 的值为:%d\n", c)

	c = a - b
	fmt.Println()
	fmt.Printf("第二行 -c 的值为：%d\n", c)
	fmt.Println()

	c = a * b

	fmt.Printf("第三行 -c的值为:%d\n", c)
	fmt.Println()

}
