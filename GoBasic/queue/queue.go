package main

import (
	"container/list"
	"fmt"
)

//Queue implementation using in built doubly linked list
type Queue struct {
	list *list.List
}

//Enqueue adds element at back
func (q *Queue) Enqueue(value int) bool {
	q.list.PushBack(value)
	return true
}

//Dequeue removes element from front
func (q *Queue) Dequeue() bool {
	//Empty list
	if q.list.Len() == 0 {
		return false
	}

	f := q.list.Front()
	q.list.Remove(f)
	return true
}

//returns element at front. If empty returns false and -1.
func (q *Queue) Front() (bool, int) {
	//Empty list
	if q.list.Len() == 0 {
		return false, -1
	}
	f := q.list.Front()
	return true, f.Value.(int)
}

//'Back' returns element at back. If empty returns false and -1.
func (q *Queue) Back() (bool, int) {
	//Empty list
	if q.list.Len() == 0 {
		return false, -1
	}
	b := q.list.Back()
	return true, b.Value.(int)
}

func (q *Queue) IsEmpty() bool {
	//Empty list
	if q.list.Len() == 0 {
		return true
	}
	return false
}

func (q *Queue) Print() {

	//Empty list
	if q.list.Len() == 0 {
		fmt.Println("Queue is Empty")
		return
	}

	fmt.Print("Queue:  ")
	for e := q.list.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}

func main() {
	q := &Queue{list: list.New()}
	q.Enqueue(5)
	q.Enqueue(13)
	q.Enqueue(3)
	q.Print()
	if ok, val := q.Front(); ok {
		fmt.Println("Queue Front:", val)
	}
	if ok, val := q.Back(); ok {
		fmt.Println("Queue Back:", val)
	}
	if ok := q.Dequeue(); ok {
		fmt.Println("Queue Dequeued")
	}
	q.Print()
	if ok, val := q.Front(); ok {
		fmt.Println("Queue Front:", val)
	}
	if ok, val := q.Back(); ok {
		fmt.Println("Queue Back:", val)
	}

}
