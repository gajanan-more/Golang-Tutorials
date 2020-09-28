package main

import (
	"fmt"
)

//This is a simple example of where to use Pointer and how we can use it

func main() {
	x := 0

	//Here we are passing value of x
	foo(x)
	fmt.Println(x)

	//Here we are passing address of x
	fmt.Println(&x)

	bar(&x)
	fmt.Println(&x)

	fmt.Println(x)

}

//func foo receives the value of x
func foo(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}

//func bar receives that address of x
func bar(z *int) {
	fmt.Println(z)
	fmt.Println(*z)
	*z = 43
	fmt.Println(z)
	fmt.Println(*z)
}
