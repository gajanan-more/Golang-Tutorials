package main

import (
	"fmt"
)

type number struct {
	num_type string
	numb     int
}

func main() {
	// Struct is a data structure which compose together
	//of values of different data types. Also known as Aggregate Data type

	no1 := number{
		num_type: "odd",
		numb:     1,
	}

	no2 := number{
		num_type: "even",
		numb:     2,
	}

	fmt.Println(no1)
	fmt.Println(no2)

	fmt.Println(no1.num_type)
	fmt.Println(no2.numb)

}
