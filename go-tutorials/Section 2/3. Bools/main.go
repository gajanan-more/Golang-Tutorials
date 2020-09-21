package main

import (
	"fmt"
)

func main() {
	var s bool

	a := 101
	b := 23

	//When we don't assign any values to the bool variable, by default the value is "false"

	fmt.Println(s)

	s = true

	fmt.Println(s)

	//You will get the output as per whether your statement is true or false
	fmt.Println(a == b)

	fmt.Println(a >= b)

}
