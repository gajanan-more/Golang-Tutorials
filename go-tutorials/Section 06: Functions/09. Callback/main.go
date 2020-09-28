package main

import (
	"fmt"
)

func main() {
	xi := []int{1, 2, 4, 5, 6, 7, 8, 9, 10, 3}
	add := sum(xi...)
	fmt.Println("Sum of all numbers", add)

	addEven := even(sum, xi...)
	fmt.Println("Sum of all numbers", addEven)
}

func sum(x ...int) int {
	total := 0

	for _, v := range x {
		total += v
	}
	return total
}

func even(f func(x ...int) int, slice ...int) int {
	var evenNum []int

	for _, v := range slice {
		if v%2 == 0 {
			evenNum = append(evenNum, v)
		}
	}

	return f(evenNum...)
}
