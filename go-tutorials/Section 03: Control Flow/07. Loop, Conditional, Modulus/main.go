//A small example where we can use loop and conditions together.
package main

import (
	"fmt"
)

func main() {
	//Print numbers divisible by 6 from 1 to 100

	for i := 1; i <= 100; i++ {
		if i%6 == 0 {
			fmt.Println(i)
		}
	}
}
