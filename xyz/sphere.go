package xyz

type Sphere struct {
	Center Point3D
	Radius float64
}

func (s Sphere) Accept(v Visitor3D) {
	v.VisitSphere(s)
}
