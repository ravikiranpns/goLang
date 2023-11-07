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
