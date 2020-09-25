package main

import (
	"fmt"
)

func main() {
	x := bar()

	fmt.Printf("Type of x is %T\n", x)

	i := x()
	fmt.Println(i)
	fmt.Printf("\nType of i is %T\n\n", i)

	//Something complex. Try to figure out what is happening here
	fmt.Println("Something complex")
	fmt.Println("", bar()())
}

func bar() func() int {
	return func() int {
		return 31
	}
}
