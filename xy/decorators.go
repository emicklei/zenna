package xy

type Translate struct {
	Translation Point
	Shape       Geometric
}

func Translated(t Point, s Geometric) Translate { return Translate{Translation: t, Shape: s} }

func (t Translate) Accept(visitor Visitor) {
	visitor.VisitTranslate(t)
}

type Scale struct {
	Scale float64
	Shape Geometric
}

func (s Scale) Accept(visitor Visitor) {
	visitor.VisitScale(s)
}

type Rotate struct {
	Angle float64 // degrees rotating counter-clock
	Shape Geometric
}

func Rotated(a float64, s Geometric) Rotate { return Rotate{Angle: a, Shape: s} }

func (r Rotate) Accept(visitor Visitor) {
	visitor.VisitRotate(r)
}

type Style struct {
	Style string
	Shape Geometric
}

func Styled(st string, sh Geometric) Style { return Style{Style: st, Shape: sh} }

func (s Style) Accept(visitor Visitor) {
	visitor.VisitStyle(s)
}
