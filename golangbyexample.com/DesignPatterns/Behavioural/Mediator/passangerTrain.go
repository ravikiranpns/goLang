package main

import "fmt"

type passanger struct {
	stationMgr mediator
}

func (p *passanger) requestArrival() {
	if p.stationMgr.canLand(p) {
		fmt.Println("Passenger: Landing")
	} else {
		fmt.Println("Passenger: Waiting")
	}
}

func (p *passanger) departure() {
	p.stationMgr.notifyFree()
	fmt.Println("Passenger: Leaving")
}

func (p *passanger) permitArrival() {
	fmt.Println("Passenger: Arrival Permitted - Landing")
}
