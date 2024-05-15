package main

import (
	"fmt"
	"sync"
)

/*
	This code has a race condition as many goroutines are accessing a shared variable 'counter'. Using mutex fixes the race condition as each goroutine locks the section
	and unlocks it only after updating the variable
*/

var counter = 0

const noOfGoroutines int = 1000

func main() {
	wg := new(sync.WaitGroup)
	mutex := new(sync.Mutex)

	for i := 0; i < noOfGoroutines; i++ {
		wg.Add(1)
		go incrementCounter(wg, mutex)
	}
	wg.Wait()
	fmt.Println(counter)
}

func incrementCounter(wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	mutex.Lock()
	defer mutex.Unlock()
	v := counter
	v++
	counter = v
}
