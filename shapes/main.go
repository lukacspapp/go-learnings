package main

import "fmt"

type area interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func main() {

	tr := triangle{base: 10, height: 10}
	s := square{sideLength: 20}

	print(tr)
	print(s)

}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.base * t.height
}

func print(a area) {
	fmt.Println(a.getArea())
}
