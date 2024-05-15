package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

/* This program shows how cancelling parent contexts will cancel child contexts because Parent context calls the cancelfunc of the child contexts in recursion.

ctx is the parent context and ctx2 is the child context. cancel is the cancelfunc of ctx and cancel2 is the cancelfunc of child context ctx2.
When 'cancel' is called due to error being thrown from worker1, we can see that parent context 'ctx' also calls the cancelfunc 'cancel2' of child context ctx2 first
(since 'worker2 cancelled' is printed first in the output).
Hence parent contexts call the cancel func of child contexts recursively when they are cancelled.
*/

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		err := worker1(ctx)
		if err != nil {
			cancel()
		}
	}()

	ctx2, cancel2 := context.WithCancel(ctx)
	go func() {
		err := worker2(ctx2)
		if err != nil {
			cancel2()
		}
	}()

	worker3(ctx)
}

func worker1(ctx context.Context) error {
	time.Sleep(150 * time.Millisecond)
	return errors.New("Failed")
}

func worker2(ctx2 context.Context) error {
	for {
		select {
		case <-time.After(1000 * time.Millisecond):
			fmt.Println("worker2 done")
			return nil
		case <-ctx2.Done():
			fmt.Println("Worker2 cancelled")
			return nil
		}
	}
}

func worker3(ctx context.Context) {
	for {
		select {
		case <-time.After(1500 * time.Millisecond):
			fmt.Println("worker3 done")
			return
		case <-ctx.Done():
			fmt.Println("Worker3 cancelled")
			return
		}
	}
}
