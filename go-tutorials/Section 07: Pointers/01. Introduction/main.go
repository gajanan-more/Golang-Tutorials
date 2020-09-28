package main

import (
	"fmt"
)

//Pointers points to the address where the value get stored
//'&' gives you the address
//'*' gives you the value stored at a particular address

func main() {
	a := 42

	//To find out t where the value 42 is stored:
	fmt.Println(&a)

	//Type of a
	fmt.Printf("\n%T", a)

	//Type of &a
	fmt.Printf("\n%T\n", &a)

	//We can store the address of a variable in another variable
	c := &a
	var b *int = &a

	fmt.Println(c)
	fmt.Println(b)

	//To get the value stored at a particular address
	fmt.Println(*b)

	//Another way to do the above thing
	fmt.Println(*&a)

	//Update value at particular address
	*b = 43
	fmt.Println(a)
}
