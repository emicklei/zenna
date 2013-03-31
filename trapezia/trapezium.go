package main

import (
	. "github.com/emicklei/zenna/xy"
	"math"
)

type Trapezium struct {
	Style  string
	Angle  int // degrees
	Length int
}

func (t Trapezium) Accept(visitor Visitor) {
	b := math.Sqrt(3) / 4 * float64(t.Length)
	angleRad := float64(t.Angle) * math.Pi / 180
	c := b / math.Tan(angleRad)
	x := []float64{0, c, float64(t.Length) - c, float64(t.Length)}
	y := []float64{0, -b, -b, 0}
	// TODO optimize
	points := []Point{}
	for i := 0; i < 4; i++ {
		points = append(points, Point{x[i], y[i]})
	}
	s := StyleWith{Polygon{points}, t.Style}
	s.Accept(visitor)
}

func (t Trapezium) Delta() Point {
	b := math.Sqrt(3) / 4 * float64(t.Length)
	return Point{b / 2, -b}
}
