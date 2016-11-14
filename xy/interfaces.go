package xy

type Visitor interface {
	VisitTranslate(t Translate)
	VisitScale(s Scale)
	VisitRotate(r Rotate)
	VisitComposite(c Composite)
	VisitPolygon(p Polygon)
	VisitStyle(s Style)
	VisitLineSegment(l LineSegment)
	VisitCircle(c Circle)
	VisitRectangle(r Rectangle)
	VisitRoundedRectangle(r RoundedRectangle)
}

type Geometric interface {
	Accept(visitor Visitor)
}
