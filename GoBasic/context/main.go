package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

/*
Demo to show how multiple goroutines are attached to a context and how they shut down once the context is canceled after a timeout.

1) Context 'ctx' is created with a timeout of 50ms which is the total time for all the goroutines to start and complete the tasks.
(1 second is the total time for program to run and shut down after tasks are completed in 50ms and cleanup of resources is done by shutting down goroutines and returning.)
2) Goroutines are started and context is attached to each goroutine.
3) Each goroutine performs the task, sleeps for 2ms and again performs the task until 50ms is completed.
4) Finally when the context gets canceled, each goroutine shuts down (since each goroutine is listening for the context to be canceled, when the context is canceled
and done method returns a channel, each goroutine returns and shuts down.)
5) Finally at the end of 1 second which is the total time for the entire program to run, num of goroutines is again back to 1 as all the 5 goroutines has shut down.

Hence context has been successfuly used to create goroutines for a process and shut the goroutines down once the process is completed.
*/
const numOfGoroutines int = 5

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()
	for i := 0; i < numOfGoroutines; i++ {
		worker(ctx, i+1)
	}
	fmt.Println("Num of goroutines = ", runtime.NumGoroutine())
	time.Sleep(1 * time.Second)
	fmt.Println("Final Num of goroutines = ", runtime.NumGoroutine())
}

func worker(ctx context.Context, id int) {
	n := 0
	go func() {
		for {
			select {
			case _, ok := <-ctx.Done():
				fmt.Println(" ok = ", ok)
				fmt.Println("Worker ", id, " shutting down.")
				return
			default:
				n++
				fmt.Println("Worker ", id, " = task ", n, " completed")
				time.Sleep(2 * time.Millisecond)
			}
		}
	}()
}
