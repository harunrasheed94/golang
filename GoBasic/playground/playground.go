package main

import "fmt"

/*
fan out fan in design pattern implementation

Fan-out is the process of splitting and assigning a task to multiple go routines where each goroutine is a worker and works on the subtask separately.
Each goroutine computes and writes its result to its own channel.

Fan-in is the process of merging results from channels of these different goroutines into a single channel.

Thus by using the fanout-fanin concurrent design pattern, you can split the task into multiple subtasks and assign each subtask to a goroutine (Fan-out) which
writes the results into its own channel. Finally results from different channels are combined into a single channel (Fanin). (Fanout-fanin can also be called
as split and merge.)


*/

/*
Following code is a demo of the fanout-fanin design pattern.
Task: Calculate sum of squares of all numbers in a slice
SubTask (to be performed by each goroutine worker): Calculate square of a number

1. 'Fanout' method fans out all the tasks among worker goroutines.

2. 'main' uses the 'FanOut' method to distribute the tasks among the worker (goroutine) functions.

3. Each worker function  will calculate the square of a number.

4. 'fanOut' will return the results channel array to 'DemoFanOutFanIn' which will send it to 'fanIn' method.

5. 'fanIn' will add the results from all the results channel and finally write it to a single channel 'finalResultChannel' which will be read by main function. ('FanIn' function uses goroutines to read from each worker goroutine channel)
('finalMergedResult' will have the results from all the individual channels.)

6. 'Main' will print the final value from 'finalResultChannel' channel.
*/
const noOfGoRoutines = 5

func main() {

	arr1 := []int{10, 21, 11, 2, 15, 13, 18, 4, 8, 9}

	resultChannels := FanOut(arr1)

	finalResultChannel := FanIn(len(arr1), resultChannels)

	fmt.Println("Sum of squares of all numbers is ", <-finalResultChannel)
}

// Fan out all the tasks among worker goroutines
func FanOut(arr1 []int) []chan int {
	var resultChannels []chan int
	taskCh := make(chan int)
	for i := 0; i < noOfGoRoutines; i++ {
		ch := make(chan int)
		resultChannels = append(resultChannels, ch)
		go Worker(taskCh, ch)
	}

	go func() {
		defer close(taskCh)
		for _, val := range arr1 {
			taskCh <- val
		}
	}()

	return resultChannels
}

// Worker that calculates square of a number
func Worker(taskCh <-chan int, resultCh chan<- int) {
	defer close(resultCh)
	for {
		val, ok := <-taskCh
		if !ok {
			return
		}
		resultCh <- val * val
	}
}

// Merge all the results from all individual workers
func FanIn(n int, resultChannels []chan int) chan int {

	var finalResultChannel chan int
	finalResultChannel = make(chan int)

	sum := 0
	mergeCh := make(chan int, n)
	merge := func(resultCh chan int) {
		for {
			val, ok := <-resultCh
			if !ok {
				return
			}
			mergeCh <- val
		}
	}

	for _, ch := range resultChannels {
		go merge(ch)
	}

	go func() {
		for i := 0; i < cap(mergeCh); i++ {
			val, ok := <-mergeCh
			if !ok {
				break
			}
			sum += val
		}
		finalResultChannel <- sum
	}()
	return finalResultChannel
}
