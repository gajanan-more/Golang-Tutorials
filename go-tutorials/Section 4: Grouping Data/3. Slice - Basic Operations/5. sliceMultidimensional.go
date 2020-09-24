package main

import (
	"fmt"
)

func main() {
	//single dimensioal slice

	even := []int{2, 4, 6, 8, 10}
	fmt.Println(even)

	odd := []int{1, 3, 5, 7, 9}
	fmt.Println(odd)

	//Multi-dimensional slice
	numbers := [][]int{even, odd}
	fmt.Println(numbers)
}
