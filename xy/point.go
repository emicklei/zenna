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
	math.Sqrt(float64((p.x * p.x) + (p.y * p.y)))
}

func (p Point) Dot(Point q) float64 {
	return (p.x * q.x) + (p.y * q.y)
}

func (p Point) DistanceTo(Point q) float64 {
	dx := q.x - p.x
	dy := q.y - p.x
	return math.Sqrt((dx * dx) + (dy * dy))
}

func (p Point) AngleTo(Point q) float64 {
	return math.Atan2(q.y-p.y, q.x-p.x)
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
