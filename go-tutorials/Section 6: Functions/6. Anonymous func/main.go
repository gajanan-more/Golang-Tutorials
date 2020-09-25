package main

import (
	"fmt"
)

func main() {
	foo()

	//Anonymous function doesn't have any name. To execute them,
	//you have to add () at closning }

	func() {
		fmt.Println("In Anonymous function")
	}()

	//Anonymous function with parameters
	para := 43
	func(x int) {
		fmt.Println("I have received", x, "parameter")
	}(para)
}

func foo() {
	fmt.Println("In foo")
}
