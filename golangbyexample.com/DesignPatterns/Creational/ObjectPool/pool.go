package main

import (
	"fmt"
	"sync"
)

type pool struct {
	idle     []iPoolObject
	active   []iPoolObject
	capacity int
	mlock    *sync.Mutex
}

// InitPool Initialize the pool
func initPool(poolObjects []iPoolObject) (*pool, error) {
	if len(poolObjects) == 0 {
		return nil, fmt.Errorf("Cannot create a pool of 0 length")
	}
	active := make([]iPoolObject, 0)
	pool := &pool{
		idle:     poolObjects,
		active:   active,
		capacity: len(poolObjects),
		mlock:    new(sync.Mutex),
	}
	return pool, nil
}
