package main

import "fmt"

// Declaring a own type with underlying type int
type own_type int

// Creating a VARIABLE of new TYPE with the IDENTIFIER "x" using the "VAR" keyword
var x own_type

func main() {
	fmt.Println(x)

	// Printing out the type of the variable "x"
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
