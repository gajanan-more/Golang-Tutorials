package main

import (
	"fmt"
)

func main() {
	x := 13

	if x == 0 {
		fmt.Println("Please enter number greater than 0")
	} else if x%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}
}
