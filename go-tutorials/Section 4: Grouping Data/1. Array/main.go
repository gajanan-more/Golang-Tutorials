package main

import (
	"fmt"
)

func main() {
	//Array: It is a building block in go
	//Syntax: var var_name [size]TYPE

	var x [5]int
	var z [5]string

	//If we don't assign any value, then it assigns the by default value
	//which get assign to the variable of that type.

	fmt.Println(x)
	fmt.Println(z)

	//Assign a value to a particular array index

	x[3] = 12
	fmt.Println(x[3])

	z[3] = "Hello World"
	fmt.Println(z[3])

	//To get the length of the array
	fmt.Println(len(x))
}
