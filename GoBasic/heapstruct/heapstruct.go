package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	Value    string
	Priority int
	Index    int
}

type MyHeap []*Item

func (m MyHeap) Len() int { return len(m) }
func (m MyHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
	m[i].Index = i
	m[j].Index = j
}

//Descending order since its max heap
func (m MyHeap) Less(i, j int) bool {
	return m[i].Priority > m[j].Priority
}

func (m *MyHeap) Push(x any) {
	n := len(*m)
	item := x.(*Item)
	item.Index = n
	*m = append(*m, item)
}

func (m *MyHeap) Pop() any {
	old := *m
	n := len(*m)
	item := old[n-1]
	old[n-1] = nil  //avoid memory leak
	item.Index = -1 // for safety
	*m = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (m *MyHeap) update(item *Item, val string, priority int) {
	item.Value = val
	item.Priority = priority
	heap.Fix(m, item.Index)
}

func main() {
	// Some items and their priorities.
	itemsMap := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	i := 0
	items := make([]*Item, len(itemsMap))
	for itemName, priority := range itemsMap {
		item := &Item{
			Value:    itemName,
			Priority: priority,
			Index:    i,
		}
		items[i] = item
		i++
	}

	maxHeap := MyHeap(items)
	heap.Init(&maxHeap)
	grape := &Item{
		Value:    "grapes",
		Priority: 1,
	}
	heap.Push(&maxHeap, grape)

	maxHeap.update(grape, "grapes", 5)
	fmt.Println("Elements ordered by priority")
	for maxHeap.Len() > 0 {
		i := heap.Pop(&maxHeap).(*Item)
		fmt.Println(i.Value, "  ", i.Priority)
	}
}
