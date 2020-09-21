package main

import (
	"fmt"
)

const (
	//Start at 0
	a = iota
	b
	c
	d
	_
	f
	g
)

const (
	//Start at 1
	Sunday = iota + 23
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	January = iota + 1
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func main() {
	fmt.Printf("Type of a variable is: %T and it's value is: %d\n", a, a)

	fmt.Printf("Type of b variable is: %T and it's value is: %d\n", b, b)

	fmt.Printf("Type of c variable is: %T and it's value is: %d\n", c, c)

	fmt.Printf("Type of d variable is: %T and it's value is: %d\n", d, d)

	fmt.Printf("Type of f variable is: %T and it's value is: %d\n", f, f)

	fmt.Printf("Type of g variable is: %T and it's value is: %d\n", g, g)

	fmt.Printf("\nType of Sunday variable is: %T and it's value is: %d\n", Sunday, Sunday)

	fmt.Printf("Type of Monday variable is: %T and it's value is: %d\n", Monday, Monday)

	fmt.Printf("Type of Tuesday variable is: %T and it's value is: %d\n", Tuesday, Tuesday)

	fmt.Printf("Type of Wednesday variable is: %T and it's value is: %d\n", Wednesday, Wednesday)

	fmt.Printf("Type of Thursday variable is: %T and it's value is: %d\n", Thursday, Thursday)

	fmt.Printf("Type of Friday variable is: %T and it's value is: %d\n", Friday, Friday)

	fmt.Printf("Type of Saturday variable is: %T and it's value is: %d\n", Saturday, Saturday)

	//Try to print out the months
}
