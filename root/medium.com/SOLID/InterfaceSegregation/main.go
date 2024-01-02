package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float32
	volume() float32
}

type square struct {
	sideLen float32
}

func (s square) area() float32 {
	return s.sideLen * s.sideLen
}

func (s square) volume() float32 {
	return 0
}

type cube struct {
	sideLen float32
}

func (c cube) area() float32 {
	return float32(math.Pow(float64(c.sideLen), 2))
}

func (c cube) volume() float32 {
	return float32(math.Pow(float64(c.sideLen), 3))
}

func areaSum(shapes ...shape) float32 {
	var sum float32
	for _, s := range shapes {
		sum += s.area()
	}

	return sum
}

func areaVolumeSum(shapes ...shape) float32 {
	var sum float32

	for _, s := range shapes {
		sum += s.area() + s.volume()
	}
	return sum
}

func main() {
	sqr := square{
		sideLen: 5,
	}

	cub := cube{
		sideLen: 5,
	}

	fmt.Println("Area sum : ", areaSum(sqr, cub))
	fmt.Println("Area and Volume Sum : ", areaVolumeSum(sqr, cub))
}
