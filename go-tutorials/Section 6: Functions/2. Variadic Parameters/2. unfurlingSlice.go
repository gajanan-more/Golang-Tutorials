package main

import (
	"fmt"
)

func main() {
	slice1 := []int{2, 3, 3, 4, 5, 5, 12}

	//If you want to share the slice as a variadic parameter,
	//you have to follow, following syntax

	sum := foo(slice1...)

	fmt.Println("Sum is ", sum)
}

//Unfurling the Slice

func foo(x ...int) int {
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("At index ", i, ", sum is ", sum)
	}
	return sum
}
