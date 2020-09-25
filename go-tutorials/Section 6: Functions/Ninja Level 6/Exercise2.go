package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	totalfromfoo := foo(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	totalfrombar := bar(x...)

	fmt.Println("Sum from foo is: ", totalfromfoo)

	fmt.Println("Sum from bar is: ", totalfrombar)

}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(y ...int) int {
	sum := 0
	for _, v := range y {
		sum += v
	}
	return sum
}
