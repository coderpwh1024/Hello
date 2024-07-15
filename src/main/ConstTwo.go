package main

import "fmt"

func main() {

	const (
		a1 = iota
		b1
		c1
		d1 = "ha"
		e1
		f1 = 100
		g1
		h1 = iota
		i1
	)

	fmt.Println(a1, b1, c1, d1, e1, f1, g1, h1, i1)

}
