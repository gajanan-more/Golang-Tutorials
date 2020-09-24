package main

import (
	"fmt"
)

func main() {

	x := []int{1, 3, 5, 7, 9}

	//It will print whole slice
	fmt.Println(x)

	//It will print value at a particular index of a slice
	fmt.Println(x[2])

	//To slice a slice, we use ":" operator

	//To print whole slice using :
	fmt.Println(x[:])

	//To print values between particular indices
	fmt.Println(x[1:3])

	//To print values from start to a particular index
	fmt.Println(x[:3])

	//To print values from a particular index to end
	fmt.Println(x[2:])

	//When we use "x:y" to slice a slice, it includes
	//value from index "x" but excludes value at index "y"
}
