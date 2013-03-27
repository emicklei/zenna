package xy

type TranslateBy struct {
	Shape       Geometric
	Translation Point
}

func (t TranslateBy) Accept(visitor Visitor) {
	visitor.VisitTranslateBy(t)
}

type ScaleBy struct {
	Shape Geometric
	Scale float64
}

func (s ScaleBy) Accept(visitor Visitor) {
	visitor.VisitScaleBy(s)
}

type RotateBy struct {
	Shape Geometric
	Angle float64
}

func (r RotateBy) Accept(visitor Visitor) {
	visitor.VisitRotateBy(r)
}

type StyleWith struct {
	Shape Geometric
	Style string
}

func (s StyleWith) Accept(visitor Visitor) {
	visitor.VisitStyleWith(s)
}
