package main

import (
	"fmt"
)

var g int
var func1 func()

func main() {

	f := func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}

	f()
	fmt.Printf("%T\n", f)

	fmt.Println(g)
	fmt.Printf("%T\n", g)

	func1 = f
	func1()
	fmt.Printf("this is func1 %T\n", func1)

	fmt.Println("done")
}
