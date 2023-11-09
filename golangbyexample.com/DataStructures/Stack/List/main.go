package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	lstack *list.List
}

func (s Stack) Push(value int) {
	s.lstack.PushFront(value)
}

func (s Stack) Pop() error {
	if s.lstack.Len() > 0 {
		ele := s.lstack.Front()
		s.lstack.Remove(ele)
	}
	return fmt.Errorf("Pop Error: Stack is empty")
}

func (s Stack) Front() (int, error) {
	if s.lstack.Len() > 0 {
		if val, ok := s.lstack.Front().Value.(int); ok {
			return val, nil
		}
		return 0, fmt.Errorf("peep Error: Stack Datatype is incorrect")
	}
	return 0, fmt.Errorf("peep Error: Stack is empty")
}

func (s Stack) Size() int {
	return s.lstack.Len()
}

func (s Stack) Empty() bool {
	return s.lstack.Len() == 0
}

func main() {
	TestLStack := &Stack{
		lstack: list.New(),
	}
	fmt.Printf("Push: 1\n")
	TestLStack.Push(1)
	fmt.Printf("Push: 2\n")
	TestLStack.Push(2)
	fmt.Printf("Size: %d\n", TestLStack.Size())
}
