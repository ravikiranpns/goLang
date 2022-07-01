package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float32
}

type square struct {
	sideLen float32
}

type calculator struct {
	total float32
}

func (c calculator) sumAreas(shapes ...interface{}) float32 {
	var sum float32

	for _, shape := range shapes {
		switch shape.(type) {
		case circle:
			r := shape.(circle).radius
			sum += math.Pi * r * r

		case square:
			l := shape.(square).sideLen
			sum += l * l
		}
	}
	return sum
}

func main() {
	c := circle{radius: 5}
	s := square{sideLen: 2}
	calc := calculator{}
	fmt.Println("Total of areas : ", calc.sumAreas(c, s))
}
