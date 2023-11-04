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
	return fmt.Errorf("Pop Error: Queue is empty")
}
