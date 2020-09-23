package main

import (
	"fmt"
)

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	x := 4
	fmt.Printf("X: %d\t\t\t%b\n", x, x)

	//Bit shifting towards right
	y := x >> 1
	fmt.Printf("Y: %d\t\t\t%b\n", y, y)

	//Bit shifting towards left
	z := x << 1
	fmt.Printf("Z: %d\t\t\t%b\n", z, z)

	w := x << 10
	fmt.Printf("Z: %d\t\t%b\n", w, w)

	fmt.Printf("KB: %d\t\t%b\n", kb, kb)
	fmt.Printf("MB: %d\t\t%b\n", mb, mb)
	fmt.Printf("GB: %d\t\t%b\n", gb, gb)
}
