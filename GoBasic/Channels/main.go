package main

import (
	"MyGoRoutinesProject/channels"
)

func main() {
	/*
		Channel c1 is used to receive a message from the goroutine 'channelRoutine' and print it. 'go control' doesn't go to the next statement from line 14
		until the message is received and goes to line 15 only after the message is received in line 14.
	*/
	// c1 := make(chan []int)

	// go channelRoutine(c1)

	//RECEIVE MESSAGE from Channel and print
	// arr := <-c1
	// fmt.Println(arr)

	//channels.DemoForLoopinChannel()
	//channels.DemoUnidirectionalChannel()
	//channels.DemoSelectinChannel()
	channels.DemoFanOutFanIn()
}

// func channelRoutine(c chan []int) {
// 	arr := []int{1, 2, 3, 4, 5}
// 	//SEND MESSAGE to CHANNEL to be passed to 'main'.
// 	c <- arr
// 	fmt.Println("Done writing the message in the channel")
// }
