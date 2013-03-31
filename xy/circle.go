package xy

import "fmt"

type Circle struct {
	Center                         Point
	Radius, StartAngle, SweepAngle float64
}

// Circ is a convenient Circle constructor
func Circ(center Point) {

}

func (c Circle) VertexAtFraction(fraction float64) Point {
	return P(c.Radius, 0).RotatedBy(c.SweepAngle*fraction + c.StartAngle).Plus(c.Center)
}

func (c Circle) Accept(visitor Visitor) {
	visitor.VisitCircle(c)
}

func (c Circle) String() string {
	return fmt.Sprintf("Circ( %v, %v )", c.Center, c.Radius)
}
