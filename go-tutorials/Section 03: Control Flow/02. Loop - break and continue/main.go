package main

import (
	"fmt"
)

func main() {
	x := 0
	for {
		x++
		//If the value of x should be less than 15, otherwise it will break the loop.
		if x > 15 {
			break
		}

		//If the value of x is completely divisible by 2 then only it will continue to print
		//the value of x otherwise it will skip the print statement.
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)

	}

}
