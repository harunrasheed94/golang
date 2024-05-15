package main

import "fmt"

// interface in golang can have ONLY METHODS and NOT FIELDS
type calculate interface {
	Area() float64
}

// Circle and Triangle structs implement calculate interface by implementing its Area() method
type Circle struct {
	radius float64
}

type Triangle struct {
	base   float64
	height float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius
}

func (t Triangle) Area() float64 {
	return 0.2 * t.base * t.height
}

func main() {

	/*

			Note:
	 		DYNAMIC POLYMORPHISM CANNOT be ACHIEVED in GOLANG using EMBEDDED STRUCT and it can be ACHIEVED ONLY using INTERFACES.
	 		EMBEDDED STRUCTS is USED to ACHEIVE normal INHERITANCE and INTERFACES are used to ACHIEVE DYNAMIC POLYMORPHISM in GOLANG.
			Therefore,
			DYNAMIC POLYMORPHISM = INTERFACES,
			EMBEDDED STRUCTS = INHERITANCE

			https://www.geeksforgeeks.org/interfaces-in-golang/
			https://medium.com/rungo/interfaces-in-go-ab1601159b3a#:~:text=Since%20the%20interface%20is%20a,s%20of%20type%20interface%20Shape%20.
	*/

	var c calculate //Creating a nil interface c of type 'calculate'

	//assign a struct of type Circle and Triangl to c (DYNAMIC POLYMORPHISM).
	c = Circle{
		radius: 5,
	}

	fmt.Println("Area of circle is ", c.Area())

	c = Triangle{
		base:   5,
		height: 10,
	}

	fmt.Println("Area of the triangle is ", c.Area())

	/*empty interface demo.
	When an interface has zero methods, it is called an empty interface.
	This is represented by interface{}. Since the empty interface has zero methods, all types implement this interface implicitly and hence by
	DYNAMIC POLYMORPHISM ALL THE TYPES can BE ASSIGNED TO the EMPTY INTERFACE i. */

	showingEmptyDemo("hi")
	showingEmptyDemo(5)
	showingEmptyDemo(c)
}

func showingEmptyDemo(i interface{}) {
	fmt.Printf("Type of passed parameter is '%T' ", i)
	fmt.Println()
}
