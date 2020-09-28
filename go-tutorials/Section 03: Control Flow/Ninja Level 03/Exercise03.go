package main

import (
	"fmt"
)

//ASCII code for Uppercase alphabets start from 65 to 90.
//Rune code point is nothing but variable with int32 type

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%d\n", i)
		for j := 0; j <= 2; j++ {
			fmt.Printf("\t%#U\n", i)
		}

	}
}
