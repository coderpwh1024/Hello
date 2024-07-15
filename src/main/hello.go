package main

import (
	"fmt"
)

func main() {

	var stockcode = 123
	var enddate = "2024-07-14"
	var url = "code=%d&endDate=%s"

	fmt.Printf(url, stockcode, enddate)

	fmt.Println()
	fmt.Println("Hello World!")
	fmt.Print("This is Golang")

	fmt.Println()
	// 语言变量
	var a string = "Runoob"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d int
	fmt.Println(d)

	var e bool
	fmt.Println(e)

}
