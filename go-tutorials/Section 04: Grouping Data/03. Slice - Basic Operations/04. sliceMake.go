package main

import (
	"fmt"
)

func main() {
	//Slice has the underlying array.
	//syntax: make([]Type, length, capacity)
	//length is lenght of array where as
	//capacity is size of underlying array

	x := make([]int, 10, 10)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 12
	x[9] = 13

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	//Here we are appending value at the 11th position (index 10th)
	//of the slice. But the underlying array size is 10. So it will
	//create one more underlying array with same size and append
	//the value at 11th position
	x = append(x, 123)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
