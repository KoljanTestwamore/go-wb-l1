package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func NewPoint(x int, y int) Point {
	return Point{
		x:x,
		y:y,
	}
}

func (p1 *Point) distance(p2 *Point) float64 {
	return math.Sqrt(math.Pow(float64(p1.x) - float64(p2.x), 2) + math.Pow(float64(p1.y) - float64(p2.y), 2))
}

func main() {
	a := NewPoint(0, 0)
	b := NewPoint(1, 1)

	fmt.Printf("%v", a.distance(&b))
}