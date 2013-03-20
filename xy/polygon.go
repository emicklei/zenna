package xy

type Polygon struct {
	Points []Point
}

func (p Polygon) Accept(visitor Visitor2D) {
	visitor.VisitPolygon(p)
}
