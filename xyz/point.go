package xyz

import "fmt"

type Point3D struct {
	X, Y, Z float64
}

func P(x, y, z float64) Point3D {
	return Point3D{x, y, z}
}

func (p Point3D) Plus(o Point3D) Point3D {
	return Point3D{p.X + o.X, p.Y + o.Y, p.Z + o.Z}
}

func (p Point3D) String() string {
	return fmt.Sprintf("P(%v,%v,%v)", p.X, p.Y, p.Z)
}
