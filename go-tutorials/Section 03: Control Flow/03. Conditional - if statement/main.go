package main

import (
	"fmt"
)

func main() {
	//This is the simple program which will give you the basic idea of
	//how if statements work
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}
	//"!true" means "not true" i.e. false using "!" operator
	if !true {
		fmt.Println("003")
	}
	if !false {
		fmt.Println("004")
	}

	//You can initialize a variable with scope limited for that particular if statement
	//and if you try to access that variable it will throw error.
	if x := 42; x == 42 {
		fmt.Println("You have initialized variable having scope for this if statement")
	}

	//This statement will throw error "31:14: undefined: x"
	fmt.Println(x)

}
