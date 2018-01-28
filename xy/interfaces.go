package xy

type Visitor interface {
	VisitTranslate(t Translate)
	VisitScale(s Scale)
	VisitRotate(r Rotate)
	VisitGroup(c Group)
	VisitPolygon(p Polygon)
	VisitStyle(s Style)
	VisitLineSegment(l LineSegment)
	VisitCircle(c Circle)
	VisitRectangle(r Rectangle)
	VisitRoundedRectangle(r RoundedRectangle)
	VisitEllipse(e Ellipse)
	VisitText(t Text)
	VisitImage(i Image)
}

type Geometric interface {
	Accept(visitor Visitor)
}
