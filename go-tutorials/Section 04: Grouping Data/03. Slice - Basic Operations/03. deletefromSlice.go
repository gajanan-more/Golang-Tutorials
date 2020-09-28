package main

import (
	"fmt"
)

func main() {
	x := []int{1, 3, 5, 7, 9}
	fmt.Println(x)

	x = append(x, 11, 13, 15, 17)
	fmt.Println(x)

	y := []int{21, 23, 25, 27, 29}

	//It will append a values from slice y at the end of slice x
	x = append(x, y...)
	fmt.Println(x)

	//To delete something from the slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
