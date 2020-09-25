package main

import (
	"fmt"
)

//Syntax: func (r receiver) identifier(parameters) (return(s)) { ...code... }

func foo() {
	fmt.Println("You are in foo")
}

func main() {
	fmt.Println("Hello, playground")

	foo()

	bar("Gajanan")

	sum := add(4, 5)
	fmt.Println(sum)

	x, y := returnMulti("James", "Bond")
	fmt.Println(x)
	fmt.Println(y)
}

func bar(s string) {
	fmt.Println("Hello from bar,", s)
}

func add(a int, b int) int {
	return (a + b)
}

//Syntax: func identifier(parameters) (returns) { ...code... }
func returnMulti(str1 string, str2 string) (string, bool) {
	a := fmt.Sprint(str1, " ", str2, `, says "I am Bond, James Bond"`)
	b := false

	return a, b
}
