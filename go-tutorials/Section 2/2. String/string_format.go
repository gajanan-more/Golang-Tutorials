package main

import (
	"fmt"
)

func main() {
	var s string
	s = "l"
	fmt.Println(s)

	bytes_s := []byte(s)

	fmt.Println(bytes_s)

	n := bytes_s[0]

	//In Decimal Format
	fmt.Println(n)

	//In Binary Format
	fmt.Printf("%b", n)

	//In Hexadecimal Format
	fmt.Printf("\n%#X", n)

}
