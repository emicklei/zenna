package main

import (
	"math"
	"os"

	"github.com/ajstarks/svgo"
	. "github.com/emicklei/zenna/svgf"
	. "github.com/emicklei/zenna/xy"
)

func main() {
	canvas := svg.New(os.Stdout)
	canvas.Start(1000, 800)
	canvas.Title("Mozara")
	canvas.Translate(500, 400)
	canvas.Rotate(90)
	canvas.Scale(1)

	vis := NewSVGPainter("", canvas)

	st := "stroke-width:2px;fill:red;stroke:black"
	//st := "stroke-width:1px;fill:#4b5cD7;stroke:#ffb800"

	moz := Mozara{40, 20}
	g := new(Composite)
	m1 := Style{st, Polygon{Mirrored(moz.Mozara_1(), AxisX(0.0))}}
	addToGroup(m1, g)

	m2 := Style{st, Polygon{Mirrored(moz.Mozara_2(), AxisX(0.0))}}
	addToGroup(m2, g)

	m3 := Style{st, Polygon{Mirrored(moz.Mozara_3(), AxisX(0.0))}}
	addToGroup(m3, g)

	m4 := Style{st, Polygon{Mirrored(moz.Mozara_4(), AxisX(0.0))}}
	addToGroup(m4, g)

	m5 := Style{st, Polygon{Mirrored(moz.Mozara_5(), AxisX(0.0))}}
	addToGroup(m5, g)

	m6 := Style{st, Polygon{Mirrored(moz.Mozara_6(), AxisX(0.0))}}
	addToGroup(m6, g)

	m7 := Style{st, Polygon{Mirrored(moz.Mozara_7(), AxisX(0.0))}}
	addToGroup(m7, g)

	g.Accept(vis)

	Translate{Point{moz.Width(), moz.Height()}, g}.Accept(vis)
	Translate{Point{0, moz.Height()}, g}.Accept(vis)
	Translate{Point{0, -moz.Height()}, g}.Accept(vis)
	Translate{Point{moz.Width(), -moz.Height()}, g}.Accept(vis)
	Translate{Point{moz.Width(), 0}, g}.Accept(vis)
	Translate{Point{-moz.Width(), 0}, g}.Accept(vis)
	Translate{Point{-moz.Width(), -moz.Height()}, g}.Accept(vis)
	Translate{Point{-moz.Width(), moz.Height()}, g}.Accept(vis)

	canvas.Gend() // scale
	canvas.Gend() // translate
	canvas.Gend() // rotate
	canvas.End()
}

func addToGroup(s Geometric, g *Composite) {
	g.Add(s)
	g.Add(Rotate{90, s})
	g.Add(Rotate{180, s})
	g.Add(Rotate{270, s})
}

var tan22d5 = math.Tan(22.5 * math.Pi / 180)
var cos22d5 = math.Cos(22.5 * math.Pi / 180)
var tan67d5 = math.Tan(67.5 * math.Pi / 180)
var sin45 = math.Sin(45 * math.Pi / 180)

type Mozara struct {
	InnerRadius float64
	StrokeWidth float64
}

func (m Mozara) Width() float64 {
	s := m.InnerRadius
	t := m.StrokeWidth
	return (2*s + t) / tan22d5
}

func (m Mozara) Height() float64 {
	return m.Width()
}

func (m Mozara) Mozara_1() []Point {
	s := m.InnerRadius
	t := m.StrokeWidth

	xa := s * (tan67d5 - 1) / (tan67d5 - tan22d5)
	ya := tan67d5 * (xa - s)

	xb := (s*(tan67d5-1) - t) / (tan67d5 - tan22d5)
	yb := tan67d5 * (xb - s)

	xc := 0.0
	yc := -s - t

	xh := 0.0
	yh := -s

	xf := -2 * (s + t) / (1 + tan22d5)
	yf := xf + s + t

	v := t * cos22d5 / sin45
	xe := xf
	ye := yf - v

	xg := (2*s + t) / (-1 - tan22d5)
	yg := -tan22d5*xg - s

	xd := (ye + s + t - xe) / (-1 - tan22d5)
	yd := ye + xd - xe

	return []Point{
		Point{xa, ya}, Point{xb, yb}, Point{xc, yc}, Point{xd, yd},
		Point{xe, ye}, Point{xf, yf}, Point{xg, yg}, Point{xh, yh},
	}
}

func (m Mozara) Mozara_2() []Point {
	s := m.InnerRadius
	t := m.StrokeWidth

	xa2 := -2 * (s + t) / (1 + tan22d5)
	ya2 := -tan22d5*xa2 - s

	xb2 := xa2
	yb2 := xa2

	xc2 := xa2 - (cos22d5 * t)
	yc2 := yb2

	xd2 := xc2
	yd2 := -tan22d5*xd2 - s - t

	xe2 := (2*s + t) / (-2 * tan22d5)
	ye2 := -tan22d5*xe2 - s - t

	xf2 := -s / tan22d5
	yf2 := 0.0

	return []Point{
		Point{xa2, ya2}, Point{xb2, yb2}, Point{xc2, yc2}, Point{xd2, yd2},
		Point{xe2, ye2}, Point{xf2, yf2},
	}
}

func (m Mozara) Mozara_3() []Point {
	s := m.InnerRadius
	t := m.StrokeWidth

	xa3 := -2 * (s + t) / (1 + tan22d5)
	ya3 := -1 * (xa3 + s + t)

	xb3 := xa3
	yb3 := tan22d5*xb3 + s

	xc3 := (2*s + t) / (-2 * tan22d5)
	yc3 := -tan22d5*xc3 - s - t

	xd3 := xc3
	yd3 := yc3 + t

	xe3 := -2*(s+t)/(1+tan22d5) - (cos22d5 * t)
	ye3 := tan22d5*xe3 + s + t

	xf3 := xe3
	yf3 := -1 * (xe3 + s + t)

	return []Point{
		Point{xa3, ya3}, Point{xb3, yb3}, Point{xc3, yc3}, Point{xd3, yd3},
		Point{xe3, ye3}, Point{xf3, yf3},
	}
}

func (m Mozara) Mozara_4() []Point {
	s := m.InnerRadius
	t := m.StrokeWidth

	xa4 := (tan67d5*(s+t) - s - t) / (tan22d5 - tan67d5)
	ya4 := tan22d5*xa4 + s + t

	xb4 := (tan67d5*(s+t) - s) / (tan22d5 - tan67d5)
	yb4 := tan22d5*xb4 + s

	xc4 := (2*s + t) / (-1 - tan22d5) // xg1
	yc4 := tan22d5*xc4 + s

	xd4 := (2*s + t) / (-2 * tan22d5) // xe2
	yd4 := -xd4 - s - t

	xe4 := xd4 + (t * cos22d5 * sin45)
	ye4 := yd4 + (t * cos22d5 * sin45)

	v := t * cos22d5 / sin45
	xf4 := (-2*(s+t) + v) / (tan22d5 + 1)
	yf4 := tan22d5*xf4 + s + t

	return []Point{
		Point{xa4, ya4}, Point{xb4, yb4}, Point{xc4, yc4}, Point{xd4, yd4},
		Point{xe4, ye4}, Point{xf4, yf4},
	}
}

func (m Mozara) Mozara_5() []Point {
	s := m.InnerRadius
	t := m.StrokeWidth

	v := t * cos22d5 / sin45
	xa5 := -2*(s+t)/(1+tan22d5) - (cos22d5 * t)
	ya5 := xa5 + s + t

	xb5 := xa5
	yb5 := ya5 - v

	xd5 := (2*s + t) / (-2 * tan22d5)
	yd5 := xd5 + s + t

	xc5 := xd5
	yc5 := yd5 - v

	return []Point{
		Point{xa5, ya5}, Point{xb5, yb5}, Point{xc5, yc5}, Point{xd5, yd5},
	}
}

func (m Mozara) Mozara_6() []Point {
	s := m.InnerRadius
	t := m.StrokeWidth

	v := t * cos22d5 / sin45

	xb6 := -2*(s+t)/(1+tan22d5) - (cos22d5 * t) // xa5
	yb6 := -2 * (s + t) / (1 + tan22d5)         // xa2

	xa6 := (xb6 + yb6 - s - t + v) / 2
	ya6 := xa6 + s + t - v

	xc6 := yb6 + s + t - v
	yc6 := yb6

	yd6 := yc6 - (t * cos22d5)
	xd6 := yd6 + s + t - v

	xf6 := (2*s + t) / (-2 * tan22d5) // xd5
	yf6 := xf6 + s + t - v            // yd5

	xe6 := xf6 + yf6 - yd6
	ye6 := yd6

	return []Point{
		Point{xa6, ya6}, Point{xb6, yb6}, Point{xc6, yc6}, Point{xd6, yd6},
		Point{xe6, ye6}, Point{xf6, yf6},
	}
}

func (m Mozara) Mozara_7() []Point {
	s := m.InnerRadius
	t := m.StrokeWidth

	w := cos22d5 * t

	xa7 := (-2 * (s + t) / (1 + tan22d5)) - w
	ya7 := xa7

	yc7 := (2*s + t) / (-2 * tan22d5) // xe2
	xc7 := yc7 + s + t

	xb7 := xa7 + w
	yb7 := yc7

	xd7 := -2 * (s + t) / (1 + tan22d5) // xb2
	xd7 += 6.5                          // HACK
	yd7 := xd7 - w - 6.5                // HACK

	return []Point{
		Point{xa7, ya7}, Point{xb7, yb7}, Point{xc7, yc7}, Point{xd7, yd7},
	}
}
