package main

import (
	. "github.com/emicklei/zenna/xy"
	"math"
)

type Trapezium struct {
	Angle  int // degrees
	Length int
}

func (t Trapezium) Points() []Point {
	b := math.Sqrt(3) / 4 * float64(t.Length)
	angleRad := R(t.Angle)
	c := b / math.Tan(angleRad)
	return []Point{PointZero, P(c, b), P(float64(t.Length)-c, b), P(float64(t.Length), 0)}
}

func (t Trapezium) Delta() Point {
	b := math.Sqrt(3) / 4 * float64(t.Length)
	return Point{b / 2, b}
}

func (t Trapezium) Accept(visitor Visitor) {
	Polygon{t.Points()}.Accept(visitor)
}
