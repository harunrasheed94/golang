package channels

import "fmt"

func DemoUnidirectionalChannel() {
	/* In this function channels are bidriectional and are converted to unidirectional channel only in the goroutine. It is not possible to use a channel as a unidirectional
	channel throughout the program because then the value send in a send-only channel can never be received and a value can never be send in a 'receive-only' channel.
	*/
	ch1 := make(chan []int)
	ch2 := make(chan int)

	arr1 := []int{4, 20, 57, 21}
	ch1 <- arr1
	go Sum(ch1, ch2)
	fmt.Println(<-ch2)
}

/* Using 2 unidirectional channels, a receive only channel 'receiveCh' to receive the values to sum and a send only channel 'sendResultCh' to send the calculated sum.
Using unidirectional channels makes it clear from which channel to receive values and to which channel to send values.
*/
func Sum(receiveCh <-chan []int, sendResultCh chan<- int) {
	arr := <-receiveCh

	sum := 0
	for _, val := range arr {
		sum = sum + val
	}
	sendResultCh <- sum
}
