package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
This code has a race condition as many goroutines are accessing a shared variable 'counter'.
*/
var counter = 0

func main() {
	wg := new(sync.WaitGroup)
	fmt.Println("CPU = ", runtime.NumCPU())
	const numberGoRoutine = 100
	wg.Add(numberGoRoutine)
	for i := 0; i < numberGoRoutine; i++ {
		go incrementcounter(wg)
	}
	wg.Wait()
	fmt.Println("Counter = ", counter)
}

func incrementcounter(wg *sync.WaitGroup) {
	v := counter
	runtime.Gosched()
	v++
	counter = v
	wg.Done()
}
