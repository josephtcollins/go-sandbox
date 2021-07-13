package main

import "fmt"

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

func main() {
	t1 := triangle{10, 5}
	sq1 := square{10}

	printArea(t1)
	printArea(sq1)

}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return t.height * t.base / 2
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
