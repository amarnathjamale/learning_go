package main

import (
	"fmt"
	"reflect"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func printArea(s shape) {
	fmt.Println("Area of the shape", reflect.TypeOf(s), "with values", reflect.ValueOf(s), "is", s.getArea())

}

func main() {
	triangle1 := triangle{
		2,
		5,
	}
	triangle2 := triangle{
		10,
		5,
	}
	square1 := square{2}
	square2 := square{4}
	var square3 square
	square3.sideLength = 3
	printArea(triangle1)
	printArea(triangle2)
	printArea(square1)
	printArea(square2)
	printArea(square3)
}
