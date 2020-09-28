package main

import (
	"fmt"
)

//Declare constants
const integer = 21
const floating = 3.14
const stringing = "Hey, How you doing?"

//Another way to declare constants
const (
	a int     = 43
	b float64 = 3.124343
	c string  = "Life is hard but you keep fighting.."
)

func main() {
	const x = 10
	fmt.Printf("Type of integer variable is: %T and it's value is: %d\n", integer, integer)

	fmt.Printf("Type of floating variable is: %T and it's value is: %f\n", floating, floating)

	fmt.Printf("Type of stringing variable is: %T and it's value is: %s\n", stringing, stringing)

	fmt.Printf("Type of a variable is: %T and it's value is: %d\n", a, a)

	fmt.Printf("Type of b variable is: %T and it's value is: %f\n", b, b)

	fmt.Printf("Type of c variable is: %T and it's value is: %s\n", c, c)

	fmt.Printf("Type of x variable is: %T and it's value is: %d\n", x, x)
}
