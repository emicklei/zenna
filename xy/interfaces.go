package xy

type Visitor interface {
	VisitTranslateBy(t TranslateBy)
	VisitScaleBy(s ScaleBy)
	VisitRotateBy(r RotateBy)
	VisitComposite(c Composite)
	VisitPolygon(p Polygon)
	VisitStyleWith(s StyleWith)
	VisitLineSegment(l LineSegment)
}

type Geometric interface {
	Accept(visitor Visitor)
}
