package main

import (
	"fmt"
)

func main() {
	bar(foo)
}

func foo() {
	fmt.Println("I am from returned function")
}

func bar(f func()) {
	f()
}
