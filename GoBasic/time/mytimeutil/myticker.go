package mytimeutil

import (
	"fmt"
	"time"
)

func MyTicker() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	//execute Task every 5 seconds 5 times
	for i := 0; i < 5; i++ {
		time := <-ticker.C
		Task(time)
	}
}

func Task(t time.Time) {
	fmt.Println("Executing task at ", t.String())
}
