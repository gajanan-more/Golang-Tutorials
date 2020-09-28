package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

//Func area will receive anything of type circle as a pointer value
//It will receive only pointer values
func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

//Here we are passing a pointer value to info()
func main() {
	c := circle{31}
	info(&c)
}
