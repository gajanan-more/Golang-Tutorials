package main

import (
	"fmt"
)

func main() {
	a := 12
	b := 14
	c := 14

	g := (c == b)
	fmt.Println(g)

	h := (a <= b)
	fmt.Println(h)

	i := (a >= c)
	fmt.Println(i)

	j := (b != c)
	fmt.Println(j)

	k := (b < c)
	fmt.Println(k)

	l := (a > c)
	fmt.Println(l)
}
