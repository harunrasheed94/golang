package main

import (
	"container/list"
	"fmt"
)

//Doubly linked list using in-built 'list' struct from "container/list" package
type dll struct {
	list *list.List
}

//used to initialize a list
func (d *dll) Init() {
	d.list = list.New()
}

//Adds to front of the list
func (d *dll) AddAtFront(val int) {
	d.list.PushFront(val)
}

//Adds to tail of the list
func (d *dll) AddAtTail(val int) {
	d.list.PushBack(val)
}

//Adds to the list at a given index if the index is within the list (0 numbered index). If the index is outside the list then it is not added and false is returned.
func (d *dll) AddAtIndex(idx, val int) bool {
	if idx > d.list.Len() {
		return false
	}

	if idx == 0 {
		d.AddAtFront(val)
	} else if idx == d.list.Len() {
		d.AddAtTail(val)
	} else {
		e := d.list.Front()
		for i := 0; i < idx-1; i++ {
			e = e.Next()
		}
		d.list.InsertAfter(val, e)
	}
	return true
}

//'Find' is used to search for an element in the list. If the element exists then return true and index else return false and -1.
func (d *dll) Find(val int) (int, bool) {
	i := 0
	for e := d.list.Front(); e != nil; e = e.Next() {
		eVal := e.Value.(int)
		if eVal == val {
			return i, true
		}
		i++
	}
	return -1, false

}

//'DeleteAtFront' deletes an element at front. If the list is empty it returns false.
func (d *dll) DeleteAtFront() bool {
	if d.list.Len() > 0 {
		e := d.list.Front()
		d.list.Remove(e)
		return true
	}
	return false
}

//'DeleteAtTail' deletes an element at tail. If the list is empty it returns false.
func (d *dll) DeleteAtTail() bool {
	if d.list.Len() > 0 {
		e := d.list.Back()
		d.list.Remove(e)
		return true
	}
	return false
}

//'DeleteAtIndex' deletes an element at given index. If the list is empty or index is outside the list size it returns false.
func (d *dll) DeleteAtIndex(idx int) bool {
	if idx >= d.list.Len() {
		return false
	}

	if idx == 0 {
		d.DeleteAtFront()
	} else if idx == d.list.Len() {
		d.DeleteAtTail()
	} else {
		e := d.list.Front()
		for i := 0; i < idx; i++ {
			e = e.Next()
		}
		d.list.Remove(e)
	}
	return true
}

func (d *dll) Print() {
	if d.list.Len() == 0 {
		fmt.Println("List is empty")
	} else {
		fmt.Print("List = ")
		for e := d.list.Front(); e != nil; e = e.Next() {
			fmt.Print(e.Value.(int))
			if e.Next() != nil {
				fmt.Print(" -> ")
			}
		}
		fmt.Println()
	}
}

func main() {
	d := &dll{}
	d.Init()
	d.AddAtFront(15)
	d.Print()
	d.AddAtTail(20)
	d.Print()
	d.AddAtFront(2)
	d.Print()
	d.AddAtIndex(1, 11)
	d.Print()
	if idx, ok := d.Find(4); ok {
		fmt.Println("Element 4 found at idx = ", idx)
	} else {
		fmt.Println("Element 4 not found")
	}

	d.AddAtIndex(0, 1)
	d.Print()
	d.AddAtIndex(2, 4)
	d.Print()
	d.DeleteAtIndex(3)
	d.Print()
	d.DeleteAtIndex(0)
	d.Print()
	d.DeleteAtIndex(d.list.Len() - 1)
	d.Print()
	if idx, ok := d.Find(4); ok {
		fmt.Println("Element 4 found at idx = ", idx)
	} else {
		fmt.Println("Element 4 not found")
	}
}
