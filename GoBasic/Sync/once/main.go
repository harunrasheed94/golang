package main

import (
	"fmt"
	"sync"
)

// A global once instance
var once sync.Once

func Initialization() {
	fmt.Println("Initialized")
}

func main() {

	noOfGoroutines := 5
	var wg *sync.WaitGroup

	for i := 0; i < noOfGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(Initialization)
		}()
	}
	wg.Wait()
	fmt.Println("DONE")

}
