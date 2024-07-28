package main

import "fmt"

type Shape interface {
	area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {

	return r.width * r.height
}

type Circle struct {
	redius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.redius * c.redius
}

func main() {

	var s Shape

	s = Rectangle{
		width:  10,
		height: 5,
	}

	fmt.Printf("矩形面积:%f\n", s.area())

	fmt.Println()

	s = Circle{redius: 3}
	fmt.Printf("原型的面积:%f\n", s.area())

}
