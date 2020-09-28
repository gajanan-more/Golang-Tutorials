package main

import (
	"fmt"
)

func main() {
	//This is an anonymous struct i.e. it doesn't have any name.
	//We declared it and added the values

	agent1 := struct {
		first_name string
		last_name  string
		age        int
	}{
		first_name: "James",
		last_name:  "Bond",
		age:        32,
	}

	fmt.Println(agent1)
}
