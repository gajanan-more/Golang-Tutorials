package main

import (
	"fmt"
)

func main() {
	//Syntax: delete(map_name, "key")
	m := map[string]int{
		"Raj":     23,
		"Sheldon": 32,
		"Penny":   21,
	}

	fmt.Println(m)

	delete(m, "Sheldon")

	fmt.Println(m)

	//If you try to delete something which does not belong to
	//map, it won't show any error

	delete(m, "Sheldon")

	fmt.Println(m)
}
