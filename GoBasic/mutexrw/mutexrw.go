package main

import (
	"fmt"
	"sync"
	"time"
)

var counter = 0

const noOfWriterGoroutines = 2
const noOfReaderGoroutines = 3

/*
In below example, you can see how:
1) When writer goroutine 1 has access to lock, no other goroutine has access to the share resource 'counter'.
2) Multiple Reader goroutines have access to the lock at the same time to perform the read operation (but writer goroutine 2 has to wait).
3) Finally writer goroutine 2 got access after all 3 reader goroutines realeased the lock.
*/

func main() {
	wg := new(sync.WaitGroup)
	rwMutex := new(sync.RWMutex)
	wg.Add(noOfWriterGoroutines + noOfReaderGoroutines)
	go incrementCounter(1, wg, rwMutex) //writer go-routine 1
	for i := 0; i < noOfReaderGoroutines; i++ {
		go readCounter(i+1, wg, rwMutex) //reader go-routines
	}
	go incrementCounter(2, wg, rwMutex) //writer go-routine 2
	wg.Wait()
	fmt.Println("final counter value = ", counter)
}

func incrementCounter(id int, wg *sync.WaitGroup, rwMutex *sync.RWMutex) {
	defer wg.Done()
	rwMutex.Lock()
	defer rwMutex.Unlock()
	fmt.Println("Writer", id, " got lock")
	fmt.Println("Incrementing counter to ", counter+1)
	counter++
	time.Sleep(2 * time.Second)
	fmt.Println("Writer", id, " releasing lock")
}

func readCounter(id int, wg *sync.WaitGroup, rwMutex *sync.RWMutex) {
	defer wg.Done()
	rwMutex.RLock()         //ReadLock
	defer rwMutex.RUnlock() //ReadUnlock
	fmt.Println("Reader", id, " got lock")
	fmt.Println("counter = ", counter)
	time.Sleep(1 * time.Second)
	fmt.Println("Reader", id, " releasing lock")
}

/*
Output:

Writer 1  got lock
Incrementing counter to  1
Writer 1  releasing lock

Reader 1  got lock
counter =  1
Reader 2  got lock
counter =  1
Reader 3  got lock
counter =  1
Reader 3  releasing lock
Reader 2  releasing lock
Reader 1  releasing lock

Writer 2  got lock
Incrementing counter to  2
Writer 2  releasing lock

final counter value =  2
*/
