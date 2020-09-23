package main

import (
	"fmt"
)

func main() {
	birth_year := 1996

	for {
		if birth_year > 2020 {
			break
		}
		fmt.Println(birth_year)
		birth_year++
	}
}
