package main

import (
	"fmt"
	"math"
)

// structure for representing x y coordinates
type Point struct {
	x, y float64
}

// init func for creating new point
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// method for calculating distance between two points
func (p Point) CalculateDistance(other Point) float64 {
	dx := other.x - p.x
	dy := other.y - p.y
	distance := math.Sqrt(dx*dx + dy*dy)
	return distance
}

func main() {
	// create two points using init func
	pointA := NewPoint(2, 3)
	pointB := NewPoint(6, 8)

	// calculate distance between two points
	distance := pointA.CalculateDistance(pointB)

	fmt.Printf("Distance between points (%.2f, %.2f) and (%.2f, %.2f): %.2f\n", pointA.x, pointA.y, pointB.x, pointB.y, distance)
}
