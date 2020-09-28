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

//Here we are passing a non-pointer value to info()
func main() {
	c := circle{10}
	info(c)

	//Note: This code won't work. It will show you following error:

	//To make it correct, comment out info(c) and\
	//call c.area(). e.g. fmt.Println(c.area())
}
