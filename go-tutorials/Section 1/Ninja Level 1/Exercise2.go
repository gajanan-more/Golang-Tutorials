package main

import "fmt"

// Three variables declared having package level scope
var x int
var y string
var z bool

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

// After running this you will receive some default values compiler assigned to these identifires.
// These values are called as Zero Values.
