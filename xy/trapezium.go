package xy

import (
	"math"
)

type Trapezium struct {
	Angle  int // degrees
	Length float64
}

func (t Trapezium) Points() []Point {
	b := math.Sqrt(3) / 4 * t.Length
	angleRad := R(t.Angle)
	c := b / math.Tan(angleRad)
	return []Point{PointZero, P(c, b), P(t.Length-c, b), P(t.Length, 0)}
}

func (t Trapezium) Delta() Point {
	b := math.Sqrt(3) / 4 * t.Length
	return Point{b / 2, b}
}

func (t Trapezium) Accept(visitor Visitor) {
	Polygon{t.Points()}.Accept(visitor)
}
