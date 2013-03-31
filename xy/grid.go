package xy

type Grid struct {
	Center Point
	Width  float64
	Height float64
	Dx     float64
	Dy     float64
}

func (g Grid) Accept(visitor Visitor) {
	c := new(Composite)
	hw := g.Width / 2
	hh := g.Height / 2
	for x := g.Center.X - hw; x <= g.Center.X+hw; x += g.Dx {
		segment := LineSegment{P(x, g.Center.Y-hh), P(x, g.Center.Y+hh)}
		c.Add(segment)
	}
	for y := g.Center.Y - hh; y <= g.Center.Y+hh; y += g.Dy {
		segment := LineSegment{P(g.Center.X-hw, y), P(g.Center.X+hw, y)}
		c.Add(segment)
	}
	c.Accept(visitor)
}
