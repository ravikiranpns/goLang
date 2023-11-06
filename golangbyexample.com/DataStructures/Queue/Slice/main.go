package main

import (
	"sync"
)

type sQ struct {
	sliceQ []string
	lock   sync.RWMutex
}

func (q *sQ) EnQ(name string) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.sliceQ = append(q.sliceQ, name)
}
