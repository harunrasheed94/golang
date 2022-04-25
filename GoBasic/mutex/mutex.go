package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
This code has a race condition as many goroutines are accessing a shared variable 'counter'. Using mutex fixes the race condition as each goroutine locks the section
and unlocks it only after updating the variable
*/
var counter = 0

func main() {
	wg := new(sync.WaitGroup)
	mutex := new(sync.Mutex)
	fmt.Println("CPU = ", runtime.NumCPU())
	const numberGoRoutine = 100
	wg.Add(numberGoRoutine)
	for i := 0; i < numberGoRoutine; i++ {
		go incrementcounter(wg, mutex)
	}
	wg.Wait()
	fmt.Println("Counter = ", counter)
}

func incrementcounter(wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	v := counter
	//runtime.Gosched()
	v++
	counter = v
	wg.Done()
	mutex.Unlock()
}
