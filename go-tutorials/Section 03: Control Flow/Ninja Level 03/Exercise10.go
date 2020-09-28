package main

import (
	"fmt"
)

func main() {
	var favSport string
	favSport = "Cricket"

	switch favSport {
	case "Football":
		fmt.Println("My favourite sport is Football")
	case "Badminton", "Cricket", "Volleyball":
		fmt.Println("My favourite sport is either Badminton, Cricket or Volleyball")
	case "Tennis":
		fmt.Println("My favourite sport is Tennis")
	default:
		fmt.Println("I don't like sport")
	}
}
