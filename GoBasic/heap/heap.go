package main

import (
	"container/heap"
	"fmt"
)

/*To use functions like init, push, pop, fix etc in heap package, your concrete type 'Minheap' has to implement heap.Interface from heap package.
Since all the above mentioned functions only take a 'heap.Interface' type as the parameter.
Hence To implement heap.Interface interface, we have to implement 2 methods Push and Pop
along with 3 methods of sort.Interface i.e. Len, Swap and Less methods as heap.Interface implements sort.Interface also.

Hence 'Minheap' will implement 5 methods 'Len', 'Swap', 'Less', 'Push' and 'Pop' to implement heap.Interface.
*/

type Minheap []int

func (a Minheap) Len() int           { return len(a) }
func (a Minheap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Minheap) Less(i, j int) bool { return a[i] < a[j] }

/*use Pointer receiver for Push and Pop as they modify the slice length (and not just contents like Swap).
Both THE PUSH and POP METHODS are CALLED by inbuilt heap.push and heap.pop methods of heap package*/

func (a *Minheap) Push(x any) {
	*a = append(*a, x.(int))
}

func (a *Minheap) Pop() (x any) {
	temp := *a
	n := len(temp)
	x = temp[n-1]
	*a = temp[:n-1]
	return x
}

//My own method to get top of heap without popping out the element
func (a *Minheap) Peek() (x any) {
	x = (*a)[0]
	return x
}

func main() {
	slice := []int{15, 12, 4, 5, 13, 10}

	h := Minheap(slice)
	heap.Init(&h) //Construct min-heap using 'init' method from heap package

	fmt.Println("Pushing element ", 2, " into heap")
	heap.Push(&h, 2)
	fmt.Println("Top of min heap is = ", h.Peek())

	fmt.Println("Popping top 2 elements from min-heap")
	for i := 0; i < 2; i++ {
		fmt.Println(heap.Pop(&h))
	}

	fmt.Println("Top of min heap is = ", h.Peek())
}
