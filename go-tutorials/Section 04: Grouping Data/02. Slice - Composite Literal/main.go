package main

import (
	"fmt"
)

func main() {
	//It is advised to not to use array in Golang and
	//instead it's advised to use Slices
	//A Slice allows you to group together Values of the same Type
	//Composite literal is the way we create different values
	//these types where we compose together a bunch of values

	//Syntax: var_name := type{values}

	x := []int{1, 3, 5, 7, 9}
	fmt.Println(x)

	//We can perform all the operations we can perform on array
	fmt.Println(len(x))
	fmt.Println(x[2])

	//We can use the range to iterate through the slice

	for i, v := range x {
		fmt.Println(i, v)
	}
}
