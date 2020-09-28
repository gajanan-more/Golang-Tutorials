package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"Raj":     23,
		"Sheldon": 32,
	}

	fmt.Println(m)

	//To add new element to the map

	m["Penny"] = 21

	//We can use range in the map as well

	for i, v := range m {
		fmt.Println(i, v)
	}
}
