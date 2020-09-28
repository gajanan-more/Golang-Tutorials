package main

import (
	"fmt"
)

func main() {
	//'&&' is Logical AND operator and '||' is Logical OR operator
	// '!' is Logical Negative
	fmt.Printf("true && true\t%v\n", true && true)
	fmt.Printf("true && false\t%v\n", true && false)
	fmt.Printf("true || true\t%v\n", true || true)
	fmt.Printf("true || false\t%v\n", true || false)
	fmt.Printf("!true\t%v\n", !true)
	fmt.Printf("!false\t%v\n", !false)
}
