package xyz

type Geometric3D interface {
	Accept(visitor Visitor3D)
}

type ExtendedGeometric3D struct {
	Origin Point3D
	Extent Point3D
}

func (e ExtendedGeometric3D) Corner() Point3D {
	return e.Origin.Plus(e.Extent)
}
