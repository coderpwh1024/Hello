package main

import "fmt"

/*
**
全局变量
*/
var g int

func main() {

	var a, b int

	a = 10

	b = 20

	g = a + b

	fmt.Printf("结果: a=%d,b=%d and g=%d\n", a, b, g)

}
