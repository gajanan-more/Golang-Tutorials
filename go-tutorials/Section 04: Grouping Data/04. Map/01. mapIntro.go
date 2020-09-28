package main

import (
	"fmt"
)

func main() {
	//Map is a key-value store, and it's an unordered list
	//syntax: map[Type]Type{"Key":"Value",}

	m := map[string]int{
		"Raj":     23,
		"Sheldon": 32,
	}

	fmt.Println(m)

	//To print out the value for some key

	fmt.Println(m["Raj"])

	//If you enter something which doesn't belong to map,
	//it will show 0 or null

	fmt.Println(m["Penny"])

	//If you wanna check whether something belongs to map or not,
	//you can use "comma-okay"(, ok) ediom

	v, ok := m["Penny"]
	fmt.Println(v)
	fmt.Println(ok)

	//You can use if statement as well

	if v, ok := m["Penny"]; ok {
		fmt.Printf("We have Penny and her age is %d", v)
	}
}
