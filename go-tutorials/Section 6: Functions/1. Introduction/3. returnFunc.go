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
	sum := add(4, 5)
	fmt.Println(sum)
}

func bar(s string) {
	fmt.Println("Hello from bar,", s)
}

//Syntax: func identifier(parameters) (return) { ...code... }
func add(a int, b int) int {
	return (a + b)
}
