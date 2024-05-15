package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	st *list.List
}

func (s *Stack) Init() {
	s.st = list.New()
}

//Push the element to top of stack
func (s *Stack) Push(val int) {
	s.st.PushBack(val)
}

//Pops the last inserted element. LIFO
func (s *Stack) Pop() (int, bool) {
	if s.st.Len() == 0 {
		return -1, false
	}
	e := s.st.Back()
	val := e.Value.(int)
	s.st.Remove(e)
	return val, true
}

//Returns top of the stack i.e. Last inserted element
func (s *Stack) Top() (int, bool) {
	if s.st.Len() == 0 {
		return -1, false
	}
	e := s.st.Back()
	val := e.Value.(int)
	return val, true

}

//Returns if stack is empty
func (s *Stack) IsEmpty() bool {
	if s.st.Len() == 0 {
		return true
	}
	return false
}

func main() {
	stack := &Stack{}
	stack.Init()

	stack.Push(4)
	stack.Push(15)
	stack.Push(18)
	stack.Push(23)
	for !stack.IsEmpty() {
		if val, ok := stack.Top(); ok {
			fmt.Println("Top of stack = ", val)
		}
		if val, ok := stack.Pop(); ok {
			fmt.Println("Popped = ", val)
		}
	}

	fmt.Println("Stack is empty")
}
