package main

import (
	"fmt"
)

//Defer always runs whenever the function is going to exit

func main() {
	fmt.Println("Hello, playground")
	defer foo()
	bar()
	hello()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

func hello() {
	fmt.Println("Hello")
}
