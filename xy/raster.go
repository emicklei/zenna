package xy

import "math"

type Raster struct {
	Center Point
	Width  float64
	Height float64
	Dx     float64 // Dx >=0 , if Dx = 0 then grid contains 1 vertical line
	Dy     float64 // Dy >=0 , if Dy = 0 then grid contains 1 horizontal line
}

func (g Raster) Accept(visitor Visitor) {
	c := new(Composite)
	hc := math.Floor(g.Width / g.Dx / 2)
	vc := math.Floor(g.Height / g.Dy / 2)
	hw := g.Width / 2
	hh := g.Height / 2
	if g.Dx != 0 {
		for i := -hc; i <= hc; i++ {
			x := i*g.Dx + g.Center.X
			segment := LineSegment{P(x, g.Center.Y-hh), P(x, g.Center.Y+hh)}
			c.Add(segment)
		}
	} else {
		segment := LineSegment{P(g.Center.X, g.Center.Y-hh), P(g.Center.X, g.Center.Y+hh)}
		c.Add(segment)
	}
	if g.Dy != 0 {
		for i := -vc; i <= vc; i++ {
			y := i*g.Dy + g.Center.Y
			segment := LineSegment{P(g.Center.X-hw, y), P(g.Center.X+hw, y)}
			c.Add(segment)
		}
	} else {
		segment := LineSegment{P(g.Center.X-hw, g.Center.Y), P(g.Center.X+hw, g.Center.Y)}
		c.Add(segment)
	}
	c.Accept(visitor)
}
