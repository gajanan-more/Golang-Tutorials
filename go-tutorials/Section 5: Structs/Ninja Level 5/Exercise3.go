package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourwheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	v1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourwheel: true,
	}

	v2 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: false,
	}

	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Println(v1.doors)
	fmt.Println(v2.doors)
	fmt.Println(v1.color)
	fmt.Println(v2.color)
	fmt.Println(v1.fourwheel)
	fmt.Println(v2.luxury)
}
