package main

import (
	"fmt"
)

//UNTYPED Constants
const (
	a = 43
	b = 23
	c = 15
)

//TYPED Constants
const (
	x int     = 34
	y float64 = 12.21
	z string  = "Try harder!!!!"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}
