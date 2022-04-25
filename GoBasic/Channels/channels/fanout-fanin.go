package channels

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

/*
Fan-out is the process of splitting and assigning a task to multiple go routines where each goroutine is a worker and works on the subtask separately.
Each goroutine computes and writes its result to its own channel.
Fan-in is the process of merging results from channels of these different goroutines into a single channel.
Thus by using the fanout-fanin concurrent design pattern, you can split the task into multiple subtasks and assign each subtask to a goroutine (Fan-out) which
writes the results into its own channel. Finally results from different channels are combined into a single channel (Fanin). (Fanout-fanin can also be called
as split and merge.)
*/
type Person struct {
	name string
	age  int
}

const GoRoutines int = 25

/*Following code is a demo of the fanout-fanin design pattern. Task is to calculate the number of persons eligible to drive. Minimum there will be 100 persons and maximum there
could be 300 persons and 25 goroutines (workers) are used to accomplish this task. Additional 25 goroutines are used in 'Fanin' to merge the results of the 25 'worker' goroutines
into a single channel.

1. 'DemoFanOutFanIn' method will generate the persons array.

2. 'DemoFanOutFanIn' uses the 'fanOut' method to distribute the tasks among the 25 worker (goroutine) functions.

3. Each worker function  will check if the person is eligible to drive or not and then write it to its own channel.

4. 'fanOut' will return the results channel array to 'DemoFanOutFanIn' which will send it to 'fanIn' method.

5. 'fanIn' will merge the results from the results channel array and write it to a single channel 'finalMergedResult' and return it to 'DemoFanOutFanin'. For this 'fanIn'
will have another 25 goroutines running for 25 worker goroutines. Each of the 25 goroutines in 'fanIn' will read the result from each worker's channel and
write it into the common 'finalMergedResult' channel.
('finalMergedResult' will have the results from all the individual channels.)

6. 'DemoFanOutFanIn' will print the final merged results from 'finalMergedResult' channel.

Although the steps explained above might seem sequential (synchronous), steps 2 to 6 all happen parallely (Asynchronously) at the same time. All the goroutines will be running
parallely and doing their tasks. So, 'fanOut' will keep writing the tasks to the consume channel and each 'worker' will keep consuming the tasks and process it. At the same time
25 goroutines in 'fanIn' also will be running and waiting for the result on each individual worker's channel, and as soon as the results start coming in will write it into the
'finalMergedResult' channel. Results in the 'finalMergedResul' channel also will be consumed by the 'DemoFanOutFanIn' method at the same time. Hence all the actions
happen concurrently at the same time with number of goroutines running.
*/
func DemoFanOutFanIn() {
	startTime := time.Now()
	min := 100
	max := 300
	n := rand.Intn(max-min) + min
	persons := make([]Person, n)

	for i := 0; i < n; i++ {
		age := rand.Intn(20) + 10
		p := Person{
			name: "Person" + strconv.Itoa(i),
			age:  age,
		}
		persons[i] = p
	}

	resultChannels := fanOut(persons)

	finalMergedResult := fanIn(resultChannels...)

	noOfValidDrivers := 0

	for result := range finalMergedResult {
		fmt.Println(result)
		noOfValidDrivers++
	}

	fmt.Println("Number of valid drivers = ", noOfValidDrivers)
	fmt.Println("Time taken to complete the tasks = ", time.Since(startTime).Seconds())
}

/*'fanOut' distributes the tasks among the 25 worker (goroutine) functions by starting each of the 25 worker goroutines. It has one goroutine running to put the tasks in the consume channel.
Each task (i.e. each person from the persons array) is put in the consume channel by the one running goroutine in fanOut and closes the 'consume' channel once
all the tasks have been consumed by each of the worker goroutine and stops running and exits.
*/
func fanOut(persons []Person) []<-chan string {
	consumeChannel := make(chan Person)
	resultChannels := make([]<-chan string, GoRoutines)
	for i := 0; i < GoRoutines; i++ {
		resultChannels[i] = worker(consumeChannel, i)
	}

	go func() {
		for _, person := range persons {
			consumeChannel <- person
		}
		close(consumeChannel)
	}()

	return resultChannels

}

/*Each worker has one goroutine to check if the person is eligible to drive or not. Once it checks if the person is eligible to drive or not, it writes the result to its own
channel and goes on to consume the next task (i.e. next person from the consume channel) after a millisecond delay. It finally shuts down when there are no more tasks to be consumed.
*/
func worker(consume <-chan Person, id int) <-chan string {
	resultChannel := make(chan string)

	go func() {
		for {
			p, ok := <-consume
			if !ok {
				break
			}
			if p.age > 18 {
				resultChannel <- p.name + " can drive by worker-" + strconv.Itoa(id)
			}
			time.Sleep(1 * time.Millisecond)

		}
		close(resultChannel)
	}()
	return resultChannel

}

/*'fanIn' is used to merge the results from each of the 'worker' goroutine's channel and write it into a common channel.
'fanIn' spawns as many goroutines as worker goroutines. Each spawned goroutine merges the result from each worker's channel to the final channel. Once each spawned goroutine
has completely consumed from its respective worker's channel (and the worker's channel has closed), it exits.
Waitgroups are used to ensure that 'fanIn' method exits only after each of the goroutine merging the results has finished execution.
*/
func fanIn(resultChannels ...<-chan string) <-chan string {
	finalMergedResults := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(len(resultChannels))

	multiplex := func(resultChannel <-chan string) {
		for {
			resStr, ok := <-resultChannel
			if !ok {
				break
			}
			finalMergedResults <- resStr
		}
		wg.Done()
	}

	for _, resultChannel := range resultChannels { //spawning as many goroutines as workers to merge the result from each worker's channel to the final channel ('finalMergedResults').
		go multiplex(resultChannel)
	}
	go func() {
		wg.Wait()
		close(finalMergedResults)
	}()

	return finalMergedResults
}
