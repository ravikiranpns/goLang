package main

import "fmt"

type node struct {
	data string
	prev *node
	next *node
}

type dll struct {
	len  int
	head *node
	tail *node
}

func initdll() *dll {
	return &dll{}
}

func (d *dll) AddFront(data string) {
	newNode := &node{
		data: data,
	}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	}
	d.len++
	return
}

func (d *dll) AddEnd(data string) {
	newNode := &node{
		data: data,
	}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		curr := d.head
		for curr.next != nil {
			curr = curr.next
		}
		newNode.prev = curr
		curr.next = newNode
		d.tail = newNode
	}
	d.len++
	return
}

func (d *dll) TraverseFwd() error {
	if d.head == nil {
		return fmt.Errorf("traverseError: List empty")
	}
	curr := d.head
	for curr != nil {
		fmt.Printf("value = %v, prev = %v, next = %v\n", curr.data, curr.prev, curr.next)
		curr = curr.next
	}
	fmt.Println()
	return nil
}

func (d *dll) TraverseRev() error {
	if d.head == nil {
		return fmt.Errorf("traverseError: List empty")
	}
	curr := d.tail
	for curr != nil {
		fmt.Printf("value = %v, prev = %v, next = %v\n", curr.data, curr.prev, curr.next)
		curr = curr.prev
	}
	fmt.Println()
	return nil
}

func (d *dll) Size() int {
	return d.len
}
