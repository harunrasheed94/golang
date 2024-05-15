package main

import (
	"fmt"
)

/*

INHERITANCE in GO using EMBEDDED STRUCTS (COMPOSITION)
 In this example, TWO CHILD STRUCTS (Marvel and DC) will USE the PARENTs (BASE = Comic's) METHOD to get the comic universe.

 By EMBEDDING PARENT STRUCT in the CHILDREN STRUCTS, Children will inherit all the properties and methods of the PARENT STRUCT and hence can use them both.

 Note:
 DYNAMIC POLYMORPHISM CANNOT be ACHIEVED in GOLANG using EMBEDDED STRUCT and it can be ACHIEVED ONLY using INTERFACES.
 EMBEDDED STRUCTS is USED to ACHEIVE normal INHERITANCE and INTERFACES are used to ACHIEVE DYNAMIC POLYMORPHISM in GOLANG.
 Therefore,
		DYNAMIC POLYMORPHISM = INTERFACES,
		EMBEDDED STRUCTS = INHERITANCE

GOLANG ACHIEVES INHERITANCE THROUGH COMPOSITION.
*/

// declaring a struct
type Comic struct {

	// declaring struct variable
	Universe string
}

// function to return the
// universe of the comic
func (comic Comic) comicUniverse() string {

	// returns comic universe
	return comic.Universe
}

// declaring a struct
type Marvel struct {

	// anonymous field,
	// this is composition where
	// the struct is embedded
	Comic
}

// Declaring a struct
type DC struct {

	// anonymous field
	Comic
}

// main function
func main() {

	// creating an instance
	c1 := Marvel{

		// child struct can directly
		// access base struct variables
		Comic: Comic{
			Universe: "MCU",
		},
	}

	// child struct can directly
	// access base struct methods

	// printing base method using child struct
	fmt.Println("Universe is:", c1.comicUniverse())

	c2 := DC{
		Comic: Comic{
			Universe: "DC",
		},
	}

	// printing base method using child
	fmt.Println("Universe is:", c2.comicUniverse())
}
