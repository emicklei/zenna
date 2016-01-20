package xy

type Visitor interface {
	VisitTranslateOn(t TranslateOn)
	VisitScaleOn(s ScaleOn)
	VisitRotateOn(r RotateOn)
	VisitComposite(c Composite)
	VisitPolygon(p Polygon)
	VisitStyleOn(s StyleOn)
	VisitLineSegment(l LineSegment)
	VisitCircle(c Circle)
	VisitRectangle(r Rectangle)
	VisitRoundedRectangle(r RoundedRectangle)
}

type Geometric interface {
	Accept(visitor Visitor)
}
