package main

import (
	"fmt"
)

// Declared a string type variable with global scope. Strings are immutable.
var str1 string

func main() {
	// Declared a string type variable with local scope
	var s string
	s = "Hello World"

	str1 = "Life is really good"

	fmt.Printf("Type of variable s is %T and value inside the variable is '%s' \n", s, s)

	// A string is a sequence of bytes. bytes_str1 will have the bytes in decimal format of whatever we have in str1. Cross check it with ASCII codes.
	bytes_str1 := []byte(str1)

	// Type of bytes_str1
	fmt.Printf("%T\n", bytes_str1)
	fmt.Println("\n", bytes_str1)

	// This is the UTF8 code point for string s
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	// This is in the Hexadecimal format. Use "X" to feel the difference.
	for i, v := range s {
		fmt.Printf("\nAt index position %d we have character in hex format %#x", i, v)
	}

}
