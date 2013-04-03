package xyz

import "fmt"

type cube struct {
	ExtendedGeometric3D
}

func Cube(origin, extent Point3D) cube {
	return cube{ExtendedGeometric3D{origin, extent}}
}

func (c cube) String() string {
	return fmt.Sprintf("Cube(%v, %v)", c.Origin, c.Extent)
}

//func (c cube) FacesDo(block func(f Face)) {}
