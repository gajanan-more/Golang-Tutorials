package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	x = append(x, 12)

	fmt.Println(x)

	x = append(x, 13, 14, 15)

	fmt.Println(x)

	y := []int{21, 22, 23, 24, 25}

	x = append(x, y...)

	fmt.Println(x)

}
