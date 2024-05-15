package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

type Data struct {
	DateTime  string `json:"datetime"`
	Value     string `json:"value"`
	Partition string `json:"partition"`
}

type Queue struct {
	Data    []*Data
	rwMutex *sync.RWMutex
	n       int
	Front   int
	Rear    int
}

func (q *Queue) Init(size int) {
	q.Data = make([]*Data, size)
	q.rwMutex = &sync.RWMutex{}
	q.n = size
	q.Front = -1
	q.Rear = -1
}

func (q *Queue) Enqueue(d *Data) bool {
	q.rwMutex.Lock()
	defer q.rwMutex.Unlock()
	if (q.Rear+1)%q.n == q.Front {
		return false
	}

	q.Rear = (q.Rear + 1) % q.n
	q.Data[q.Rear] = d
	if q.Front == -1 {
		q.Front = 0
	}

	return true
}

func (q *Queue) Dequeue() *Data {
	q.rwMutex.Lock()
	defer q.rwMutex.Unlock()
	if q.Front == -1 && q.Rear == -1 {
		return nil
	}
	d := q.Data[q.Front]

	if q.Front == q.Rear {
		q.Front = -1
		q.Rear = -1
	} else {
		q.Front = (q.Front + 1) % q.n
	}
	return d
}

func main() {
	q := &Queue{}
	q.Init(30)

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file. Error = ", err.Error())
		return
	}
	defer file.Close()
	wg := &sync.WaitGroup{}
	wg.Add(2)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	ch := make(chan int)
	go Producer(q, wg, fileScanner, ch)
	go Consumer(q, wg, ch)
	wg.Wait()
}

/*
Producer produces to the queue and then lets the consumer know it has produced by sending a message on channel which the consumer is listening to.
*/
func Producer(q *Queue, wg *sync.WaitGroup, fileScanner *bufio.Scanner, ch chan int) {
	defer wg.Done()
	defer close(ch)
	for fileScanner.Scan() {
		str := fileScanner.Text()
		d := &Data{}
		err := json.Unmarshal([]byte(str), d)
		if err != nil {
			fmt.Println(err)
			return
		}

		//Producer will enqueue the data. If the queue is full then it will sleep for 1 millisecond and try again to enqueue until it succeeds.
		for {
			if enqueued := q.Enqueue(d); enqueued {
				break
			}
			time.Sleep(1 * time.Millisecond)
		}
		ch <- 1
	}
}

/*Consumer consumes from the queue to which the producer produces.
Producer writes to the channel after it produces every record which the Consumer listens to.
Consumer starts consuming only after producer has produced.*/
func Consumer(q *Queue, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for {
		_, ok := <-ch
		data := q.Dequeue()
		if !ok && data == nil { //if channel is closed and there is no data in queue, then return
			return
		}
		if data != nil {
			fmt.Println(*data)
		}
	}

}
