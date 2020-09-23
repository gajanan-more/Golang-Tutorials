package main

import (
	"fmt"
)

func main() {
	//Here are some small examples of int and float and their conversion. You can try by yourself and play with it to understand it better.
	var point int8 = 13

	var circle int32

	circle = int32(point)

	fmt.Println(circle)

	fmt.Printf("Point of type:    %T\n", point)
	fmt.Printf("Circle of type: %T\n", circle)

	var large int64 = 34

	var small int8

	small = int8(large)

	fmt.Println(small)

	large = 139
	small = int8(large)
	fmt.Println(small)

	var i int64 = 67

	var j float64 = float64(i)

	fmt.Printf("%.4f\n", j)

}
