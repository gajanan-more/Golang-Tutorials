package main

import (
	"fmt"
)

//Recursion calls the func from itself
func main() {
	fact := factorial(6)

	fmt.Println("Factorial of number is: ", fact)
}

func factorial(numb int) int {
	if numb < 1 {
		return 1
	}
	fmt.Println(numb)
	return numb * factorial(numb-1)

}

//It's not advised to use recursive because of memory utilization.
