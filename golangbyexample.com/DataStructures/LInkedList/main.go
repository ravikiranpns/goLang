package main

import (
	"fmt"
)

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

func (s *singleList) AddFront(num int) int {
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
	return 0
}

func (s *singleList) AddBack(num int) int {
	newNode := &node{
		data: num,
	}
	if s.head == nil {
		s.head = newNode
	} else {
		curr := s.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = newNode
	}
	s.len++
	return 0
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

func (s *singleList) Size() int {
	return s.len
}

func (s *singleList) RemoveFront() error {
	if s.head == nil {
		return fmt.Errorf("list is Empty")
	}
	s.head = s.head.next
	s.len--
	return nil
}

func (s *singleList) RemoveBack() error {
	if s.head == nil {
		return fmt.Errorf("removeBack: List is empty")
	}

	var prev *node
	curr := s.head
	for curr.next != nil {
		prev = curr
		curr = curr.next
	}

	if prev != nil {
		prev.next = nil
	} else {
		s.head = nil
	}
	s.len--
	return nil
}

func (s *singleList) Front() (int, error) {
	if s.head == nil {
		return 0, fmt.Errorf("single List is empty")
	}
	return s.head.data, nil
}

func main() {
	ssl := initList()

	err := ssl.Traverse()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("AddFront: 1\n")
	ssl.AddFront(1)
	fmt.Printf("AddFront: 2\n")
	ssl.AddFront(2)
	fmt.Printf("AddBack: 3\n")
	ssl.AddBack(3)

	fmt.Printf("Size: %d\n", ssl.Size())
	err = ssl.Traverse()
	if err != nil {
		fmt.Println(err.Error())
	}

}
