package main

import (
	"fmt"
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

func (q *sQ) DeQ() error {
	if len(q.sliceQ) > 0 {
		q.lock.Lock()
		defer q.lock.Unlock()
		q.sliceQ = q.sliceQ[1:]
		return nil
	}
	return fmt.Errorf("pop Error: Q is empty")
}

func (q *sQ) Front() (string, error) {
	if len(q.sliceQ) > 0 {
		q.lock.Lock()
		defer q.lock.Unlock()
		return q.sliceQ[0], nil
	}
	return "", fmt.Errorf("peep Error : Q is empty")
}

func (q *sQ) Size() int {
	return len(q.sliceQ)
}

func (q *sQ) Empty() bool {
	return len(q.sliceQ) == 0
}

func main() {
	testSliceQ := &sQ{
		sliceQ: make([]string, 0),
	}

	fmt.Printf("EnQ: A\n")
	testSliceQ.EnQ("A")
	fmt.Printf("EnQ: B\n")
	testSliceQ.EnQ("B")
	fmt.Printf("Len: %d\n", testSliceQ.Size())

}
