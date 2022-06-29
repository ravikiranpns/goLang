package main

import (
	"fmt"
	"math"
)

//type circle with its area method
type circle struct {
	radius float32
}

func (c circle) area() {
	fmt.Printf("Circle area : %f", math.Pi*c.radius*c.radius)
}

type square struct {
	sideLen float32
}

func (s square) area() {
	fmt.Printf("Square area : %f", s.sideLen*s.sideLen)
}

func main() {
	c := circle{radius: 5}
	c.area()

	s := square{sideLen: 2}
	s.area()
}
