package main

import (
	"fmt"
	"math"
)

/*
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
*/

//type circle with its area method
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

type shape interface {
	area() float32
}

type outPrinter struct{}

func (op outPrinter) toText(s shape) string {
	return fmt.Sprintf("The Area is : %f", s.area())
}

func main() {
	c := circle{radius: 3}
	c.area()

	s := square{sideLen: 2}
	s.area()

	out := outPrinter{}

	fmt.Println(out.toText(c))
	fmt.Println(out.toText(s))
}
