package xy

type Grid struct {
	Center Point
	Width  float64
	Height float64
	Dx     float64
	Dy     float64
	Style  string
}

// func (g Grid) DrawOn(canvas *svg.SVG) {
// 	svgf := &SVGF{canvas}
// 	hw := float64(g.Width) / 2
// 	hh := float64(g.Height) / 2
// 	for x := g.Center.X - hw; x < g.Center.X+hw; x += g.Dx {
// 		svgf.Line(x, g.Center.Y-hh, x, g.Center.Y+hh, g.Style)
// 	}
// }
