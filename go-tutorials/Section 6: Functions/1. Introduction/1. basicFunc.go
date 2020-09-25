package main

import (
	"fmt"
)

//Functions are all about being modular
//Modular means to take code and break it
//up into small modular chunks called modules

//Syntax: func identifier() { ...code... }

func foo() {
	fmt.Println("You are in foo")
}

func main() {
	fmt.Println("Hello, playground")
	foo()
}
