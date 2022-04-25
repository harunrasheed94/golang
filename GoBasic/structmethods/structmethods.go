package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) setPersonDetails(first string, last string, age int) person {
	p.first = first
	p.last = last
	p.age = age

	return p
}

func (p person) printDetails() {
	fmt.Println("Printing Person details")
	fmt.Println("First name is = "+p.first+" ", p.last+" ", p.age)
}

func main() {
	p := person{}
	p = p.setPersonDetails("Haroon", "Rasheed", 26)
	p.printDetails()
}
