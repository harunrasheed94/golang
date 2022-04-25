package main

import "fmt"

//interface in golang can have ONLY METHODS and NOT FIELDS
type calculate interface {
	Area() float64
}

//Circle and Triangle structs implement calculate interface by implementing its Area() method
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
		DYNAMIC POYMORPHISM:
		You are not allowed to create an instance of the interface.
		But you are allowed to create a variable of an interface type and this variable can be assigned with a concrete type value that has the methods the interface requires.
		For example:
		We CANNOT INSTANTIATE "CALCULATE" OBJECT. i.e.
		c := calculate{

		} NOT ALLOWED

		But WE CAN DECLARE a VARIABLE of "CALCULATE" type
		var c calculate
		AND THEN THIS VARIABLE "c" CAN BE ASSIGNED WITH A CONCRETE TYPE VALUE (STRUCT) SUCH as "CIRCLE" and "TRIANGLE" THAT IMPLEMENT the CALCULATE INTERFACE by IMPLEMENTING ITS METHOD(s) ("Area").
		c = Circle{
			radius: 5,
		}
		and
		c = Triangle{
			base:   5,
			height: 10,
		}

		Thus DYNAMIC POLYMORPHISM is ACHIEVED USING INTERFACES in golang. 
		When a type (i.e. struct) implements an interface, a variable of that type can also be represented as the type of an interface.
		We can confirm that by creating a nil interface c of type 'calculate' and assign a struct of type Circle and Triangle.

		Dynamic type of c is Cicle and then Triangle.
		Static Type of c is Calculate

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
