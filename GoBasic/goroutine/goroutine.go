package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup) //wg is of *sync.WaitGroup type
	wg.Add(2)
	fmt.Println("OS = ", runtime.GOOS)
	fmt.Println("Number of CPUs = ", runtime.NumCPU())

	go printAlphabets(wg) //WaitGroup should ONLY BE PASSED AS A POINTER
	go printNumbers(wg)   //WaitGroup should ONLY BE PASSED AS A POINTER
	fmt.Println("Number of Go Routines = ", runtime.NumGoroutine())
	fmt.Println("Main finished execution")
	wg.Wait()
}

func printAlphabets(wg *sync.WaitGroup) {
	for c := 97; c < 107; c++ {
		fmt.Printf(" %c ", c)
	}
	wg.Done()
}

func printNumbers(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i+1)
	}
	wg.Done()
}
