package main

import "fmt"

type own_type int

var x own_type

var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	// CONVERSION to convert the TYPE of the VALUE stored in "x" to the UNDERLYING TYPE
	// Then ASSIGN that value to "y"
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
