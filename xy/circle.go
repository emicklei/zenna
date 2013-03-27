package xy

import "fmt"

type Circle struct {
	Center                         Point
	Radius, StartAngle, SweepAngle float64
}

// Circle_ is a convenient Circle constructor
func Circle_(center Point) {

}

func (c Circle) VertexAtFraction(fraction float64) Point {
	return P(c.Radius, 0).RotatedBy(c.SweepAngle*fraction + c.StartAngle).Plus(c.Center)
}

func (c Circle) Accept(visitor Visitor) {
	visitor.VisitCircle(c)
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle_(%v,%v)", c.Center, c.Radius)
}
