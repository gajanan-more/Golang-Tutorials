package main

import (
	"fmt"
)

type person_info struct {
	first_name string
	last_name  string
	fav_flavor []string
}

func main() {
	person1 := person_info{
		first_name: "Albert",
		last_name:  "Einstein",
		fav_flavor: []string{"Butterscotch", "Mango"},
	}

	person2 := person_info{
		first_name: "Isaac",
		last_name:  "Newton",
		fav_flavor: []string{"Hazelnut", "Pineapple", "Chocolate"},
	}

	fmt.Println(person1.first_name)
	fmt.Println(person1.last_name)
	fmt.Println("Fav flavours are:")
	for i, v1 := range person1.fav_flavor {
		fmt.Println(i, v1)
	}

	fmt.Println()

	fmt.Println(person2.first_name)
	fmt.Println(person2.last_name)
	fmt.Println("Fav flavours are:")
	for i, v2 := range person1.fav_flavor {
		fmt.Println(i, v2)
	}
}
