package main

import (
	"fmt"
)

type person struct {
	fName   string
	lName   string
	favFood []string
}

func (p person) walk() string {
	return fmt.Sprintln(p.fName, "is walking")
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fouWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	// //slices composite literals
	// s := []int{1, 2, 34, 5, 6, 7, 8}
	// fmt.Println(s)

	// for i := range s {
	// 	fmt.Println(i)
	// }

	// for i, v := range s {
	// 	fmt.Println(i, v)
	// }

	// //slices map composite literals
	// m := map[string]int{
	// 	"baby":      1,
	// 	"computers": 3,
	// 	"games":     5,
	// }

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	p1 := person{
		fName:   "John",
		lName:   "Doe",
		favFood: []string{"Pizza", "Burger", "Pasta"},
	}

	fmt.Println(p1.favFood[0])
	for i, v := range p1.favFood {
		fmt.Println(i, v)
	}

	s := p1.walk()
	fmt.Println(s)

	veh := vehicle{
		doors: 4,
		color: "red",
	}

	tr := truck{
		vehicle:  veh,
		fouWheel: true,
	}

	sed := sedan{
		vehicle: veh,
		luxury:  true,
	}

}
