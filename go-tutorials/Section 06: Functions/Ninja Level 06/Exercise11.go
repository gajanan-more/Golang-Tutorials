package main

import (
	"fmt"
)

func main() {
	n := 5

	fmt.Println("Fibonacci Series for", n, " is:")

	for i := 0; i < n; i++ {
		fmt.Println(fib(i))
	}

}

func fib(f int) int {
	if f == 0 {
		return 0
	}
	if f == 1 || f == 2 {
		return 1
	} else {
		return (fib(f-1) + fib(f-2))
	}
}
