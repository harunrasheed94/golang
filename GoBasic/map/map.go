package main

import (
	"fmt"
)

var m map[string]int //ONLY DECLARING a MAP NOT ASSIGNING MEMORY for it yet.

func main() {
	/* Alternate way of declaring maps in golang
	m = map[string]int{
		"Mohammed": 1,
		"Haroon":   2,
		"Rasheed":  3,
	}*/

	/*Like in java Where you use 'new' keyword to create an empty hashmap, in golang we have to use 'make' keyword to create an empty map.
	
	'make' is used to assign memory to the map 'm' and now the map 'm' is not 'nil' (i.e. not null). 
	
	Therefore 'var m map[string]int' is used only to declare a map but not used to assign memory and the map 'm' is still nil and 'make' keyword is 
	used to assign memory to the map.

	'map' keyword in golang is == to 'new' keyword in java.
	*/

	m = make(map[string]int)
	m["Mohammed"] = 1
	m["Haroon"] = 2
	m["Rasheed"] = 3

	if val, ok := m["Haroon"]; ok {
		fmt.Println("Haroon is present and his val is = ", val)
	}

	str := "Kalil"

	m[str] = 4

	fmt.Println("\nIterating through the map")
	for v, k := range m {
		fmt.Println(v, k)
	}

	fmt.Println("\nDeleting a key in the map")
	m["hi"] = 5
	if v, ok := m["hi"]; ok {
		fmt.Println("deleting hi with cal = ", v)
		delete(m, "hi")
	}

}
