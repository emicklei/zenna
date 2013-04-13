package xy

type Visitor interface {
	VisitTranslateBy(t TranslateBy)
	VisitScaleBy(s ScaleBy)
	VisitRotateBy(r RotateBy)
	VisitComposite(c Composite)
	VisitPolygon(p Polygon)
	VisitStyleWith(s StyleWith)
	VisitLineSegment(l LineSegment)
	VisitCircle(c Circle)
	VisitRectangle(r Rectangle)
	VisitRoundedRectangle(r RoundedRectangle)
}

type Geometric interface {
	Accept(visitor Visitor)
}
