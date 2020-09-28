package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	license bool
}

func main() {

	agent1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		license: true,
	}

	agent2 := secretAgent{
		person: person{
			first: "Ethon",
			last:  "Hunt",
			age:   31,
		},
		license: false,
	}

	fmt.Println(agent1)
	fmt.Println(agent2)

	agent1.info()
}

//Syntax: func (r receiver) identifier(parameters) (returns) { ...code... }

func (s secretAgent) info() {
	fmt.Println(s.first, s.last)
}

//In the above function, the function info has access to anybody of
//type secretAgent. Here info() is available for receiver secretAgent or
//we can say we have attached info function to value of type secretAgent
