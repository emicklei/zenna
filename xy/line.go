package xy

type LineSegment struct {
	Begin Point
	End   Point
}

func (l LineSegment) Accept(visitor Visitor) {
	visitor.VisitLineSegment(l)
}
