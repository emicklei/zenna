package xy

import "fmt"

type Ellipse struct {
	Center                                   Point
	RadiusX, RadiusY, StartAngle, SweepAngle float64
}

func (c Ellipse) VertexAtFraction(fraction float64) Point {
	// TODO
	return PointZero
}

func (c Ellipse) Accept(visitor Visitor) {
	visitor.VisitEllipse(c)
}

func (c Ellipse) String() string {
	return fmt.Sprintf("Ellipse( %v, %v, %v )", c.Center, c.RadiusX, c.RadiusY)
}
