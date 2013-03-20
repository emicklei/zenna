package xy

type DotPen struct {
	delta      Point
	points     []Point
	offset     Point
	tipX, tipY float64
}

func NewDotPen(start, delta Point) DotPen {
	return DotPen{offset: start, delta: delta}
}

func (p *DotPen) Path() []Point {
	if p.points == nil {
		return []Point{p.offset}
	}
	return p.points
}

func (p *DotPen) Move(timesX, timesY float64) {
	p.tipX += timesX
	p.tipY += timesY
	next := Point{p.delta.X*p.tipX + p.offset.X, p.delta.Y*p.tipY + p.offset.Y}
	p.points = append(p.points, next)
}
