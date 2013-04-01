package xyz

type Sphere struct {
	Center Point
	Radius float64
}

func (s Sphere) Accept(v Visitor) {
	v.VisitCircle(s)
}
