package main

import (
	"fmt"
)

func foo() {
	fmt.Println("You are in foo")
}

func main() {
	fmt.Println("Hello, playground")
	foo()
	bar("Gajanan")
}

//Syntax: func identifier(parameters) { ...code... }
func bar(s string) {
	fmt.Println("Hello from bar,", s)
}
