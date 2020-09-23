package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("It will never print something")
	case true:
		fmt.Println("It will always print anything")
	}
}
