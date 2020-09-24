package main

import (
	"fmt"
)

func main() {

	//We have append function to append something to the slice
	//We can append unlimited no of values to the slice

	x := []int{1, 3, 5, 7, 9}
	fmt.Println(x)

	x = append(x, 11, 13, 15, 17)
	fmt.Println(x)

	y := []int{21, 23, 25, 27, 29}

	//It will append a values from slice y at the end of slice x
	x = append(x, y...)
	fmt.Println(x)
}
