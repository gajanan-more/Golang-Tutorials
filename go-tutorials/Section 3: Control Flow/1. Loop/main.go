package main

import (
	"fmt"
)

//There are multiple syntaxes of using for statement. We don't have while loop in Golang,
//so we can use for as while loop as well.

func main() {
	//Syntax: for init; condition; post {}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	//Syntax: for condtion {}
	x := 1
	for x < 10 {
		fmt.Printf("X = %d\n", x)
		x++
	}

	fmt.Println()

	//Nested for loop
	for x := 0; x < 10; x++ {
		fmt.Printf("The outer loop: %d\t\n", x)
		for y := 0; y < 3; y++ {
			fmt.Printf("\t\tThe inner loop: %d\n", y)
		}
	}

	fmt.Println()

	//We can use range in the for loop as well.
	s := "Hello World!!!"
	for i, v := range s {
		fmt.Printf("\nAt index position %d we have character in hex format %#x", i, v)
	}

}
