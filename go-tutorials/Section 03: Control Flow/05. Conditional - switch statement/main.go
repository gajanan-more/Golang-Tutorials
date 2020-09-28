package main

import (
	"fmt"
)

func main() {
	//This switch case will print something from the first true case
	switch {
	case false:
		fmt.Println("It will never print something")
	case (2 == 3):
		fmt.Println("The case is true and it will print 2 == 3")
	case (3 == 3):
		fmt.Println("The case is true and it will print 3 == 3")
	case (5 == 5):
		fmt.Println("The case is true and it will print 5 == 5")
	case (2 == 3):
		fmt.Println("The case is true and it will print 2 == 3")
	default:
		fmt.Println("This is the default case")
	}

	//We can use fallthrough to print the next statements as well using fallthrough
	//It won't check whether the case is true or not,
	//it will do whatever is there under next case

	switch {
	case false:
		fmt.Println("It will never print something")
	case (2 == 3):
		fmt.Println("The case is true and it will print 2 == 3")
	case (3 == 3):
		fmt.Println("The case is true and it will print 3 == 3")
		fallthrough
	case (5 == 10):
		fmt.Println("The case is true and it will print 5 == 10")
	case (2 == 3):
		fmt.Println("The case is true and it will print 2 == 3")
	default:
		fmt.Println("This is the default case")
	}

	//This is the small example of demonstrating how default works
	//how to add multiple cases under one case

	compare_var := "Live, Love, Laugh"

	switch compare_var {
	case "Date", "Time", "Live, Love":
		fmt.Println("It's true")
	case "Laugh":
		fmt.Println("Live, Love and Laugh")
	default:
		fmt.Println("This is the default case and nothing matches with compare_var")
	}

}
