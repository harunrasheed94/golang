package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Addr Address
	// if map is added then 'Person' struct DOESN'T SATISFY COMPARABLE because map can't be compared using standard comparison operators
	//Map  map[string]string
}

type Address struct {
	City  string
	State string
}

// Find is a Generic function used to find the index of an element in a slice where the type of the element should support comparison using standard comparison operators.
func Find[T comparable](elements []T, target T) int {
	idx := -1
	for i, val := range elements {
		if val == target { //Type T should be comparable using '=='
			idx = i
			break
		}
	}
	return idx
}

func main() {

	//int is comparable
	intSlice := []int{20, 19, 51, 33, 11, 71, 10}

	findInt := 19
	fmt.Println("Found ", findInt, " in index ", Find(intSlice, findInt))

	//string is comparable
	stringSlice := []string{"hi", "good", "knowledge", "on", "Golang", "generics"}
	findStr := "Golang"
	fmt.Println("Found ", findStr, " in index ", Find(stringSlice, findStr))

	//Custom type 'Person' (Struct) is comparable as the struct has ONLY COMPARABLE TYPES (like string, int)
	person1 := Person{
		Name: "Mohammed",
		Age:  23,
		Addr: Address{City: "Tampa", State: "Florida"},
	}

	person2 := Person{
		Name: "Haroon",
		Age:  26,
		Addr: Address{City: "Dallas", State: "Texas"},
	}

	person3 := Person{
		Name: "Rasheed",
		Age:  29,
		Addr: Address{City: "SF", State: "California"},
	}

	person4 := Person{
		Name: "Hari",
		Age:  28,
		Addr: Address{City: "Austin", State: "Texas"},
	}
	personSlice := []Person{person1, person2, person3, person4}
	fmt.Println("Found Haroon in index ", Find(personSlice, person2))
}

/*
Output:
Found  19  in index  1
Found  Golang  in index  4
Found Haroon in index  1
*/
