package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

/*
Hence Below example confirms the fact that Method set of Type *T can be used by both T and *T and Method Set of Type T can also be used by both Type T and *T.
Therefore a method set of Type T INCLUDES all METHODS DECLARED with BOTH TYPE T and TYPE *T as the RECEIVER

*/

func (p *Person) displayPersonPointer() {
	fmt.Println("\nPointer function")
	fmt.Println("Person ", p.Name, "'s age is ", p.Age)
}

func (p Person) displayPersonValue() {
	fmt.Println("\nValue function")
	fmt.Println("Person ", p.Name, "'s age is ", p.Age)
}

func main() {
	p1Value := Person{
		Name: "Haroon Value",
		Age:  29,
	}

	//Type T can access Method Set of Type T and *T
	p1Value.displayPersonPointer()
	p1Value.displayPersonValue()

	p2Pointer := &Person{
		Name: "Rasheed Pointer",
		Age:  29,
	}

	//Type *T can access Method Set of Type T and T
	p2Pointer.displayPersonPointer()
	p2Pointer.displayPersonValue()

}
