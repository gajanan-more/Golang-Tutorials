package main

import (
	"fmt"
)

func main() {
	foo(2, 3, 3, 4, 5, 5, 12)
	zeroVariadic()
	multParawithVariadic("Life", 1, 2, 3, 4, 4)
}

//Func with Variadic parameter takes unlimited values
//Parameter store those values it in form of Slice
//Variadic means zero or more

func foo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T", x)
}

func zeroVariadic(y ...int) {
	fmt.Println("\ninside zeroVariadic func")
	fmt.Println(y)
}

//Another parameter along with variadic parameter
func multParawithVariadic(s string, x ...int) {
	fmt.Println("\n", s)
	fmt.Println(x)
}
