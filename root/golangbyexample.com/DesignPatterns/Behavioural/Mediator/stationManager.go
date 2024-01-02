package main

import "sync"

type stMgr struct {
	/* 	canLand(train) bool
	   	notifyFree() */
	isPlatformFree bool
	lock           *sync.Mutex
	trainQue       []train
}

func NewStationManager() *stMgr {
	return &stMgr{
		isPlatformFree: true,
		lock:           &sync.Mutex{},
	}
}

func (sM *stMgr) canLand(t train) bool {
	sM.lock.Lock()
	defer sM.lock.Unlock()
	if sM.isPlatformFree {
		sM.isPlatformFree = false
		return true
	}
	sM.trainQue = append(sM.trainQue, t)
	return false
}

func (sm *stMgr) notifyFree() {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	if !sm.isPlatformFree {
		sm.isPlatformFree = true
	}
	if len(sm.trainQue) > 0 {
		firstTrainInQue := sm.trainQue[0]
		sm.trainQue = sm.trainQue[1:]
		firstTrainInQue.permitArrival()
	}
}
