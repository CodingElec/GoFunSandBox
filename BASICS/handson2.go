package main

import "fmt"

type person struct {
	fname string
	lname string
}

type sagent struct {
	person
	agentNumber int
}

func (p person) speak() {
	fmt.Println("Good morning, from", p.fname, p.lname)
}

func (sa sagent) speak() {
	fmt.Println("Good morning, from", sa.fname, sa.lname, "my badge number is", sa.agentNumber)
}
