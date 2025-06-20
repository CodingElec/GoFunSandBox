package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) calcArea() float64 {
	return s.side * s.side
}

func (c circle) calcArea() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	calcArea() float64
}

func area(sh shape) {
	fmt.Println(sh.calcArea())
}

func main() {

	// sq := square{2}
	// circ := circle{2}
	// fmt.Printf("the side of the square is %.2f, and the radius is %.2f \n", sq.side, circ.radius)
	// fmt.Printf("Area of the square: %.2f, area of the circle: %.2f\n", sq.calcArea(), circ.calcArea())
	// area(sq)
	// area(circ)

	investigado := person{
		"Iron",
		"man",
	}

	agent := sagent{
		person: person{
			fname: "Agent",
			lname: "Colton",
		},
		agentNumber: 123456,
	}

	agent2 := sagent{
		person:      investigado,
		agentNumber: 12345678,
	}

	investigado.speak()
	agent.speak()
	agent2.person.speak()
	agent2.speak()

	println("----- from the polymorph -----")

	speaks(investigado)
	speaks(agent)
	speaks(agent2.person)
	speaks(agent2)
}
