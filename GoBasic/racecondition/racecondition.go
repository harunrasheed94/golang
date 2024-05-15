package main

import (
	"fmt"
	"sync"
)

/*
This code has a race condition as many goroutines are accessing a shared variable 'counter'.
*/
var counter = 0

const noOfGoroutines = 1000

func main() {
	wg := new(sync.WaitGroup)
	for i := 0; i < noOfGoroutines; i++ {
		wg.Add(1)
		go incrementCounter(wg)
	}
	wg.Wait()
	fmt.Println(counter)
}

func incrementCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	v := counter
	v++
	counter = v
}
