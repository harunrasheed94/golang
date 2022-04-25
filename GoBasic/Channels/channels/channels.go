package channels

import (
	"fmt"
	"strconv"
)

/*'DemoForLoopinChannel' is used to show how to use 'for' loop to read multiple values from a channel
 */
func DemoForLoopinChannel() {
	ch := make(chan string)
	go sendMultipleValuesChannel(ch)

	i := 0
	for {
		val, ok := <-ch
		if !ok {
			fmt.Printf("Channel has received all the values.")
			break
		}
		fmt.Println("Val ", i, " = ", val)
		i++
	}
}

//'sendMultipleValuesChannel' is used to send multiple values inside a channel using for loop
func sendMultipleValuesChannel(ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- "val" + strconv.FormatInt(int64(i), 10)
	}
	close(ch) //VERY IMPORTANT to CLOSE THE CHANNEL so that the RECEIVER KNOWS that there are no more values left in the channel and it can break the infinite for loop.
}
