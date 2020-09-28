package main

import (
	"fmt"
)

func main() {
	//Func expression means assigning a function to a variable
	//Here f is of type function
	f := func() {
		fmt.Println("In Anonymous function")
	}

	f()

	para := 43

	//It receives the parameter of type int
	j := func(x int) {
		fmt.Println("I have received", x, "parameter")
	}

	j(para)

	fmt.Printf("Type of j is %T", j)

	//It returns the parameter of type int
	k := func() int {
		return 3
	}

	fmt.Printf("\nType of k is %T", k)
}
