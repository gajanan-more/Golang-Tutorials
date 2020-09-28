package main

import (
	"fmt"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func main() {
	s1 := square{
		length: 12,
		width:  13,
	}

	c1 := circle{
		radius: 4,
	}

	fmt.Println("Area of Square:", s1.area())

	fmt.Println("Area of Circle:", c1.area())

	info(s1)
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("Area of your shape is", s.area())
}
