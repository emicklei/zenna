package xy

type TranslateOn struct {
	Translation Point
	Shape       Geometric
}

func (t TranslateOn) Accept(visitor Visitor) {
	visitor.VisitTranslateOn(t)
}

type ScaleOn struct {
	Scale float64
	Shape Geometric
}

func (s ScaleOn) Accept(visitor Visitor) {
	visitor.VisitScaleOn(s)
}

type RotateOn struct {
	Angle float64
	Shape Geometric
}

func (r RotateOn) Accept(visitor Visitor) {
	visitor.VisitRotateOn(r)
}

type StyleOn struct {
	Style string
	Shape Geometric
}

func (s StyleOn) Accept(visitor Visitor) {
	visitor.VisitStyleOn(s)
}
