package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Gajanan",
		last:  "More",
		age:   21,
	}

	p1.speak()
}

func (p person) speak() {
	fmt.Println("My first name is", p.first, "and my last name is", p.last)
	fmt.Println("My age is", p.age)
}
