package main

import (
	"fmt"
)

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	fmt.Println(x)

	fmt.Println(y)

	z := [][]string{x, y}

	fmt.Println(z)

	for i, rang1 := range z {
		fmt.Println("record: ", i)
		for j, v := range rang1 {
			fmt.Printf("\tIndex Position: %v \t value: %v\n", j, v)
		}
	}
}
