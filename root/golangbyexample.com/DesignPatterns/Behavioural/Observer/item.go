package main

import "fmt"

type item struct {
	observerList []observer
	name         string
	inStock      bool
}

func newItem(name string) *item {
	return &item{
		name: name,
	}
}

func (i *item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.notifyAll()
}

func (i *item) register(o observer) {
	i.observerList = append(i.observerList, o)
}

func (i *item) deregister(o observer) {
	i.observerList = append(i.observerList, o)
}

func (i *item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

func removeFromObserverSlice(observerList []observer, observerToRemove observer) []observer {
	obsListLen := len(observerList)

	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[obsListLen-1], observerList[i] = observerList[i], observerList[obsListLen-1]
			return observerList[:obsListLen-1]
		}
	}
	return observerList
}
