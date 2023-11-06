package main

import (
	"container/list"
	"fmt"
)

type LQ struct {
	lq *list.List
}

func (q *LQ) Enqueue(value string) {
	q.lq.PushBack(value)
}

func (q *LQ) Dequeue() error {
	if q.lq.Len() > 0 {
		ele := q.lq.Front()
		q.lq.Remove(ele)
	}
	return fmt.Errorf("pop error: Queue is empty")
}

func (q *LQ) Front() (string, error) {
	if q.lq.Len() > 0 {
		if val, ok := q.lq.Front().Value.(string); ok {
			return val, nil
		}
		return "", fmt.Errorf("peep Error: Queue datatype is incorrect")
	}
	return "", fmt.Errorf("peep Error: Queue is empty")
}

func (q *LQ) Size() int {
	return q.lq.Len()
}

func (q *LQ) Empty() bool {
	return q.lq.Len() == 0
}

func main() {
	clq := &LQ{
		lq: list.New(),
	}

	fmt.Printf("Enqueue: A\n")
	clq.Enqueue("A")
	fmt.Printf("Enqueue: B\n")
	clq.Enqueue("B")
	fmt.Printf("Size: %d\n", clq.Size())

	for clq.Size() > 0 {
		fVal, _ := clq.Front()
		fmt.Printf("Front: %s\n", fVal)
		fmt.Printf("Dequeue: %s\n", fVal)
		clq.Dequeue()
	}
	fmt.Printf("Size: %d\n", clq.Size())
}
