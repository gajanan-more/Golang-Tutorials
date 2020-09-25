package main

import (
	"fmt"
)

func main() {
	//Simple
	s := struct {
		first_name string
		last_name  string
		age        int
	}{
		first_name: "Gajanan",
		last_name:  "More",
		age:        23,
	}

	fmt.Println(s)

	//Difficult
	p := struct {
		first string
		last  string
		sport []string
		cars  map[string]int
	}{
		first: "Sachin",
		last:  "Tendulkar",
		sport: []string{"Cricket", "Tennis"},
		cars: map[string]int{
			"Bently": 4,
			"Audi":   2,
			"BMW":    4,
		},
	}

	fmt.Println(p)
	//Try to use the range over sport and cars
}
