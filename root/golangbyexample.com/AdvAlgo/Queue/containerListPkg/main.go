/*
https://golangbyexample.com/queue-in-golang/

container/list - https://pkg.go.dev/container/list
- Package list implements a double linked list.
- To iterate over a list(where l is a *List)
	for e := l.Front(); e != nil; e = e.Next() {
		// do something with e.Value
	}
- A queue will have below operations:
	1.Enqueu
	2.Dequeue
	3.Front
	4.Size
	5.Empty
*/

package main

import (
	"container/list"
	"fmt"
)

type customQueue struct {
	queue *list.List
}

func (c *customQueue) Enqueue(value string) {
	c.queue.PushBack(value)
}

func (c *customQueue) Dequeue() error {
	if c.queue.Len() > 0 {
		ele := c.queue.Front()
		c.queue.Remove(ele)
	}
	return fmt.Errorf("Pop Error: Queue is empty")
}

func (c *customQueue) Front() (string, error) {
	if c.queue.Len() > 0 {
		if val, ok := c.queue.Front().Value.(string); ok {
			return val, nil
		}
		return "", fmt.Errorf("Peep Error: Queue Datatype is incorrect")
	}
	return "", fmt.Errorf("Peep Error: Queue is empty")
}

func (c *customQueue) Size() int {
	return c.queue.Len()
}

func (c *customQueue) Empty() bool {
	return c.queue.Len() == 0
}

func main() {
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	Q := &customQueue{
		queue: list.New(),
	}
	fmt.Printf("EnQ: A\n")
	Q.Enqueue("A")
	fmt.Printf("EnQ: B\n")
	Q.Enqueue("B")
	fmt.Printf("Size: %d\n", Q.Size())

	for Q.Size() > 0 {
		fVal, _ := Q.Front()
		fmt.Printf("Front: %s\n", fVal)
		fmt.Printf("Dequeue: %s\n", fVal)
		Q.Dequeue()
	}
	fmt.Printf("Size: %d\n", Q.Size())
}
