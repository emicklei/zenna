package xy

type Line struct {
	a, b, c, a2plusb2sqrt                       float64
	perpendicularUnitVector, parallelUnitVector Point
}

type LineSegment struct {
	Begin Point
	End   Point
}

func (l LineSegment) Accept(visitor Visitor) {
	visitor.VisitLineSegment(l)
}
