package main

import (
	"fmt"
)

func main() {
	var a int
	a = 50

	fmt.Printf("Number in decimal: %d", a)
	fmt.Printf("\nNumber in binary: %b", a)
	fmt.Printf("\nNumber in hexadecimal: %#x\n", a)

	x := a << 1

	fmt.Println("\nAfter shifting it to left")
	fmt.Printf("\nNumber in decimal: %d", x)
	fmt.Printf("\nNumber in binary: %b", x)
	fmt.Printf("\nNumber in hexadecimal: %#X", x)

}
