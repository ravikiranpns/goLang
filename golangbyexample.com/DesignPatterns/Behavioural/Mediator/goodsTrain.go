package main

import "fmt"

type goods struct {
	stationMgr mediator
}

func (g *goods) requestArrival() {
	if g.stationMgr.canLand(g) {
		fmt.Println("Goods: Landing")
	} else {
		fmt.Println("Goods: Waiting")
	}
}

func (g *goods) departure() {
	g.stationMgr.notifyFree()
	fmt.Println("Goods: Leaving")
}

func (g *goods) permitArrival() {
	fmt.Println("Goods: Permitted Arrival - Landing")
}
