package main

import (
	"fmt"
)

var x, y int
var (
	a int
	b int
)

var c, d int = 1, 2

var e, f = 123, "hello"

func main() {

	g, h := 123, "hello"
	fmt.Println(x, y, a, b, c, d, e, f, g, h)

	fmt.Println("----------------------")
	var m string = "abc"
	fmt.Println("hello world", m)
}
