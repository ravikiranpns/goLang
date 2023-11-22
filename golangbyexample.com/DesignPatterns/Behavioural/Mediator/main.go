package main

func main() {
	stationManager := NewStationManager()
	pTrain := &passanger{
		stationMgr: stationManager,
	}

	gTrain := &goods{
		stationMgr: stationManager,
	}

	pTrain.requestArrival()
	gTrain.requestArrival()
	pTrain.departure()
}
