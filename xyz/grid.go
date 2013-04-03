package xyz

type Grid3D struct {
	Center Point3D
	Delta  Point3D
}

func (g Grid3D) PointsDo(block func(p Point3D)) {

}
