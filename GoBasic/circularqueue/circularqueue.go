package main

import "fmt"

//since slice of fixed size is used circular queue is the best option as it doesn't waste any memory
type MyCircularQueue struct {
	arr   []int
	n     int
	front int
	rear  int
}

func Constructor(k int) MyCircularQueue {
	q := MyCircularQueue{}
	q.arr = make([]int, k)
	q.n = k
	q.front, q.rear = -1, -1
	return q
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	//if next rear has reached front then queue is full, so don't insert this element else insert the element
	if (this.rear+1)%this.n == this.front {
		return false
	}
	this.rear = (this.rear + 1) % this.n
	this.arr[this.rear] = value
	if this.front == -1 {
		this.front = 0
	}
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	//queue is empty
	if this.front == -1 && this.rear == -1 {
		return false
	}

	//we have reached the last element
	if this.front == this.rear {
		this.front = -1
		this.rear = -1
	} else {
		this.front = (this.front + 1) % this.n
	}

	return true
}

func (this *MyCircularQueue) Front() int {
	//queue is empty
	if this.front == -1 && this.rear == -1 {
		return -1
	}
	return this.arr[this.front]
}

func (this *MyCircularQueue) Rear() int {
	//queue is empty
	if this.front == -1 && this.rear == -1 {
		return -1
	}
	return this.arr[this.rear]
}

func (this *MyCircularQueue) IsEmpty() bool {
	//queue is empty
	if this.front == -1 && this.rear == -1 {
		return true
	}
	return false
}

func (this *MyCircularQueue) IsFull() bool {
	//if next rear position is front, then queue is full
	if (this.rear+1)%this.n == this.front {
		return true
	}
	return false
}

func main() {
	q := Constructor(10)
	q.EnQueue(5)
	q.EnQueue(10)
	q.EnQueue(99)
	q.EnQueue(18)
	fmt.Println(q.Front())
	fmt.Println(q.Rear())
	fmt.Println(q.IsFull())
	q.DeQueue()
	q.DeQueue()
	q.DeQueue()
	q.DeQueue()
	q.IsFull()
}
