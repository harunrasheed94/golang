package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

const eligibleAge int = 18

func main() {
	wg := new(sync.WaitGroup)

	p := &Person{
		Name: "Haroon",
		Age:  23,
	}

	noOfGoroutines := 2
	wg.Add(noOfGoroutines)
	go printPersonJSON(wg, p)
	go checkIfPersonEligibleForDriving(wg, p)
	wg.Wait()
	fmt.Println("Goroutines finished")
}

func printPersonJSON(wg *sync.WaitGroup, p *Person) {
	defer wg.Done()
	personBytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error in converting to json")
		return
	}

	fmt.Println("json: ", string(personBytes))
}

func checkIfPersonEligibleForDriving(wg *sync.WaitGroup, p *Person) {
	defer wg.Done()
	if p.Age < eligibleAge {
		fmt.Println("Person ", p.Name, " can't drive")
		return
	}
	fmt.Println("Person ", p.Name, " can drive")
}
