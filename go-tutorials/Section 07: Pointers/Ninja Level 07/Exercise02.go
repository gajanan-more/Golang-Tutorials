package main

import (
	"fmt"
)

type person struct {
	name    string
	address string
	age     int
}

func changeMe(p *person) {
	(*p).address = "India"

	//We can also do it in another way
	//p.address = "India"
}
func main() {
	p1 := person{
		name:    "Gajanan",
		address: "Maharashtra",
		age:     23,
	}

	fmt.Println(p1)

	changeMe(&p1)

	fmt.Println(p1)
}
