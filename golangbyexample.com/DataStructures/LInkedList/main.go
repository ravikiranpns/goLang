package main

import "fmt"

type node struct {
	data int
	next *node
}

type singleList struct {
	len  int
	head *node
}

func initList() *singleList {
	return &singleList{}
}

func (s *singleList) AddFront(num int) {
	newNode := &node{
		data: num,
	}

	if s.head == nil {
		s.head = newNode
	} else {
		newNode.next = s.head
		s.head = newNode
	}
	s.len++
	return
}

func (s *singleList) Traverse() error {
	if s.head == nil {
		return fmt.Errorf("error: List is empty")
	}

	cur := s.head
	for cur != nil {
		fmt.Println(cur.data)
		cur = cur.next
	}
	return nil
}
