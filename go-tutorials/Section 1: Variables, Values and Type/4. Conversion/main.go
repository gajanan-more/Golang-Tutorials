package main

import (
	"fmt"
)

//In golang, we can declare a variable of our own type.

type owntype int

//Here variable of my own type is "owntype" but the underlying type is int

func main() {
	var number owntype
	number = 42
	fmt.Println(number)
	fmt.Printf("Type of identifier number is %T\n", number)

	x := int(number)

	fmt.Println(x)
	fmt.Printf("Type of identifier x is %T", x)
}
