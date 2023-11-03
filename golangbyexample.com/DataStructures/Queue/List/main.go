package main

import (
	"container/list"
)

type Q struct {
	q *list.List
}

func (q Q) Enqueue(value string) {
	q.q.PushBack(value)
}
