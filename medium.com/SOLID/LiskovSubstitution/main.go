package main

import "fmt"

type transport interface {
	getName() string
}

// base entity
type vehicle struct {
	name string
}

func (c vehicle) getName() string {
	return c.name
}

type car struct {
	vehicle
	wheel int
	gates int
}

type motobyke struct {
	vehicle
	wheel int
}

type printer struct {
}

func (printer) printTransportName(p transport) {
	fmt.Println("Name: ", p.getName())
}

func main() {
	v := vehicle{name: "Maruti-Suzuki"}

	c := car{
		vehicle: vehicle{name: "Wagan-R"},
		wheel:   4,
		gates:   2,
	}

	m := motobyke{
		vehicle: vehicle{name: "Yamah-Crux"},
		wheel:   2,
	}

	p := printer{}
	p.printTransportName(v)
	p.printTransportName(c)
	p.printTransportName(m)

}
