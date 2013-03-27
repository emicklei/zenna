package xy

type Polygon struct {
	Points []Point
}

func (p Polygon) Accept(visitor Visitor) {
	visitor.VisitPolygon(p)
}
