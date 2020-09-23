package main

import (
	"fmt"
)

func main() {
	x := 23

	//A classic simple example to check
	//whether the number is odd or even using if-else statement
	if x%2 == 0 {
		fmt.Println("x is an even number")
	} else {
		fmt.Println("x is an odd number")
	}

	y := 50

	//if-else if-else
	if y == 45 {
		fmt.Println("The value of y is 45")
	} else if y > 45 {
		fmt.Println("The value of y is greater than 45")
	} else {
		fmt.Println("The value of y is smaller than 45")
	}
	//You can add as many if-else if-else as you want
}
