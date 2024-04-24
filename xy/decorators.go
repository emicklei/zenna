package xy

type Translate struct {
	Translation Point
	Shape       Geometric
}

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

func (r Rotate) Accept(visitor Visitor) {
	visitor.VisitRotate(r)
}

type Style struct {
	Style string
	Shape Geometric
}

func (s Style) Accept(visitor Visitor) {
	visitor.VisitStyle(s)
}
