package xy

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// P is a convenient Point constructor
func P(x, y float64) Point {
	return Point{x, y}
}

func (p Point) Plus(d Point) Point {
	return Point{p.X + d.X, p.Y + d.Y}
}

func (p Point) Multiply(num float64) Point {
	return Point{p.X * num, p.Y * num}
}

func (p Point) ScaledBy(scalePoint Point) Point {
	return Point{p.X * scalePoint.X, p.Y * scalePoint.Y}
}

func (p Point) Scaled(x, y float64) Point {
	return Point{p.X * x, p.Y * y}
}

// RotatedBy computes a new Point by rotating it clock-wise by @angle radians.
func (p Point) RotatedBy(angle float64) Point {
	sinVal := math.Sin(-angle)
	cosVal := math.Cos(-angle)
	return Point{
		X: cosVal*(p.X) + (sinVal * p.Y),
		Y: cosVal*(p.Y) - (sinVal * p.X),
	}
}

func FloatSlices(points []Point) (x, y []float64) {
	for _, each := range points {
		x = append(x, each.X)
		y = append(y, each.Y)
	}
	return x, y
}

func (p Point) String() string {
	return fmt.Sprintf("P(%v,%v)", p.X, p.Y)
}
