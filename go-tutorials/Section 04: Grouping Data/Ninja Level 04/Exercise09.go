package main

import (
	"fmt"
)

func main() {
	ppl := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	ppl["jakob martin"] = []string{"Life", "Money", "Laugh"}

	for i, v := range ppl {
		fmt.Println("These are the fav things of: ", i)
		for j, u := range v {
			fmt.Printf("\tAt index position %d, the fav thing is %v \n", j, u)
		}
	}
}
