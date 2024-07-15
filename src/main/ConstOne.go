package main

import "unsafe"

const (
	aa = "abc"
	bb = len(aa)
	cc = unsafe.Sizeof(aa)
)

func main() {

	println(aa, bb, cc)

	/*const LENGTH int = 10

	const WIDTH int = 5

	var area int

	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH

	fmt.Printf("面积为: %d ", area)
	fmt.Println()

	fmt.Println(a, b, c)*/

}
