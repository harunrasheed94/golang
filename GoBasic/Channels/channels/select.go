package channels

import "fmt"

func DemoSelectinChannel() {

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go EvenAndOddSend(even, odd, quit)
	EvenAndOddReceive(even, odd, quit)
	fmt.Println("Quitting.")
}

func EvenAndOddReceive(even, odd, quit <-chan int) {
	for {
		select {
		case val := <-even:
			fmt.Println("Even channel val = ", val)
		case val := <-odd:
			fmt.Println("Odd channel val = ", val)
		case val := <-quit:
			fmt.Println("Going to quit ", val)
			return
		}
	}

}

func EvenAndOddSend(even, odd, quit chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	quit <- 0
}
