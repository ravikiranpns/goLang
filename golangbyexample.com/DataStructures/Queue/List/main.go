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
