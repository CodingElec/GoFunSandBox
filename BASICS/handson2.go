package main

type person struct {
	fname string
	lname string
}

type sagent struct {
	person
	agentNumber int
}
