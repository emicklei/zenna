package xy

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

var PointZero = P(0, 0)

// P is a convenient Point constructor
func P(x, y float64) Point {
	return Point{x, y}
}

func (p Point) Length() float64 {
	return math.Sqrt(float64((p.X * p.X) + (p.Y * p.Y)))
}

func (p Point) Dot(q Point) float64 {
	return (p.X * q.X) + (p.Y * q.Y)
}

func (p Point) DistanceTo(q Point) float64 {
	dx := q.X - p.X
	dy := q.Y - p.Y
	return math.Sqrt((dx * dx) + (dy * dy))
}

func (p Point) AngleTo(q Point) float64 {
	return math.Atan2(q.Y-p.Y, q.X-p.X)
}

// Polar returns a Point on a Circle with a radius and angle in radians.
func Polar(radius float64, angle float64) Point {
	return Point{math.Cos(angle) * radius, math.Sin(angle) * radius}
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
