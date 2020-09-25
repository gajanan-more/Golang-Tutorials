package main

import (
	"fmt"
)

//Value can be of more than one type and interfaces allow that
//But when you try to print their type, it will show their original type

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
	per1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

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

	fmt.Printf("Agent1 is of type: %T\n", agent1)
	fmt.Printf("Per1 is of type: %T\n", per1)

	per1.info()
	agent2.info()

	bar(per1)
}

//Polymorphism: Function with the same name but a different type of receiver.
//That means both per1 of type person and agent1 of type secretAgent can call
//function info and it will work fine.

func (s secretAgent) info() {
	fmt.Println(s.first, s.last)
}

func (p person) info() {
	fmt.Println(p.first, p.last)
}

//Interfaces: It allows values to be of multiple types

type human interface {
	info()
}

func bar(h human) {
	fmt.Println("I was in the bar with", h)

	switch h.(type) {
	case person:
		fmt.Println("He is of type Person")
	case secretAgent:
		fmt.Println("He is of type SecretAgent")
	}
}

//Those who have access to info() is also of type human
//That means, any value of type secretAgent is also of type human. And same with person
//And both can have access to bar() because it accepts a parameter of type human

//Fun Fact: Every variable of some type is nothing but the empty interface in golang
