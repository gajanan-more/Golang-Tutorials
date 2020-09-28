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

	m := map[string]person_info{
		person1.last_name: person1,
		person2.last_name: person2,
	}

	//We can do it in two ways

	//first

	fmt.Println(m[person1.last_name].first_name)
	fmt.Println(m[person1.last_name].last_name)
	fmt.Println("Fav flavours are:")
	for i, v1 := range m[person1.last_name].fav_flavor {
		fmt.Println(i, v1)
	}

	fmt.Println()

	fmt.Println(m[person2.last_name].first_name)
	fmt.Println(m[person2.last_name].last_name)
	fmt.Println("Fav flavours are:")
	for i, v2 := range m[person2.last_name].fav_flavor {
		fmt.Println(i, v2)
	}

	//second
	fmt.Println("\nSecond way is\n")

	for x, val1 := range m {
		fmt.Println(x)
		fmt.Println(val1.first_name, val1.last_name)
		for y, val2 := range val1.fav_flavor {
			fmt.Println(y, val2)
		}
	}
}
