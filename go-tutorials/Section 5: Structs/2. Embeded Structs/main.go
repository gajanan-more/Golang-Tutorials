package main

import (
	"fmt"
)

//Inner Type
type person struct {
	first_name string
	last_name  string
	age        int
}

//Outer Type
type secretAgent struct {
	person
	license bool
}

func main() {

	agent1 := secretAgent{
		person: person{
			first_name: "James",
			last_name:  "Bond",
			age:        32,
		},
		license: true,
	}

	agent2 := secretAgent{
		person: person{
			first_name: "Ethon",
			last_name:  "Hunt",
			age:        31,
		},
		license: false,
	}

	fmt.Println(agent1)
	fmt.Println(agent2)

	//We can print first_name by two methods
	fmt.Println(agent1.person.first_name)
	fmt.Println(agent1.first_name)
}
