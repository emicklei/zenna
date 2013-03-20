package xy

type Visitor2D interface {
	VisitTranslateBy(t TranslateBy)
	VisitScaleBy(s ScaleBy)
	VisitRotateBy(r RotateBy)
	VisitShapeComposite(s ShapeComposite)
	VisitPolygon(p Polygon)
	VisitStyleWith(s StyleWith)
}

type Geometric2D interface {
	Accept(visitor Visitor2D)
}

type TranslateBy struct {
	Shape       Geometric2D
	Translation Point
}

func (t TranslateBy) Accept(visitor Visitor2D) {
	visitor.VisitTranslateBy(t)
}

type ScaleBy struct {
	Shape Geometric2D
	Scale float64
}

func (s ScaleBy) Accept(visitor Visitor2D) {
	visitor.VisitScaleBy(s)
}

type RotateBy struct {
	Shape Geometric2D
	Angle float64
}

func (r RotateBy) Accept(visitor Visitor2D) {
	visitor.VisitRotateBy(r)
}

type StyleWith struct {
	Shape Geometric2D
	Style string
}

func (s StyleWith) Accept(visitor Visitor2D) {
	visitor.VisitStyleWith(s)
}
