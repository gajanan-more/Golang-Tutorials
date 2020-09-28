package main

import (
	"fmt"
)

func main() {
	sum := foo(4, 5)
	fmt.Println("Sum is: ", sum)

	i, j := bar()
	fmt.Println("I got integer", i, "and string", j, "from bar")
}

func foo(x int, y int) int {
	return x + y
}

func bar() (int, string) {
	return 8, "Life"
}
