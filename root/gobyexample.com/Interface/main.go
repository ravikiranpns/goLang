package main

import (
	"fmt"
	"math"
)

// https://gobyexample.com/interfaces

// https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

type geometry interface {
	area() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	//fmt.Println(g.are
}

func main() {
	fmt.Println("The interface")
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
