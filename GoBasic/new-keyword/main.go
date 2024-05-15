package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// Create a new instance of the Person struct using new.
	p := new(Person)

	// Initialize the fields of the Person struct.
	p.FirstName = "John"
	p.LastName = "Doe"
	p.Age = 30

	// Print the person's information.
	fmt.Println("First Name:", p.FirstName)
	fmt.Println("Last Name:", p.LastName)
	fmt.Println("Age:", p.Age)
}
