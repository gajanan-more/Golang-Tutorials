package main

import (
	"fmt"
)

func main() {
	t := foo()
	fmt.Printf("\nMy type is %T", t)
	t()
}

func foo() func() {
	return func() {
		fmt.Println("I am from returned function")
	}
}
