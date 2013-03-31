package xy

type Pen struct {
	delta      Point
	recorder   PenRecorder
	offset     Point
	tipX, tipY float64
}

func NewPen(start, delta Point) Pen {
	return Pen{offset: start, delta: delta, recorder: new(PointRecorder)}
}

func (p Pen) Path() []Point {
	return p.recorder.Path()
}

func (p *Pen) Move(timesX, timesY float64) {
	p.tipX += timesX
	p.tipY += timesY
	next := Point{p.delta.X*p.tipX + p.offset.X, p.delta.Y*p.tipY + p.offset.Y}
	p.recorder.MovedTo(p, next)
}

type PenRecorder interface {
	MovedTo(pen *Pen, next Point)
	Path() []Point
}

type PointRecorder struct {
	points []Point
}

func (r *PointRecorder) MovedTo(pen *Pen, next Point) {
	r.points = append(r.points, next)
}

func (r *PointRecorder) Path() []Point {
	if r.points == nil {
		return []Point{}
	}
	return r.points
}
