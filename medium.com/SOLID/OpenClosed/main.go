package main

import (
	"math"
)

/* type circle struct {
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
} */

type shape interface {
	area() float32
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

type square struct {
	sideLen float32
}

func (s square) area() float32 {
	return s.sideLen * s.sideLen
}

type triangle struct {
	height float32
	base   float32
}

func (t triangle) area() float32 {
	return t.base * t.height / 2
}
