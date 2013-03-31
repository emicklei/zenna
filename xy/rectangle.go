package xy

import "fmt"

type Rectangle struct {
	Origin, Extent Point
}

// Rect is used to evaluate the literal representation or Rectangle.
func Rect(origin, extent Point) Rectangle {
	return Rectangle{origin, extent}
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rect( %v, %v )", r.Origin, r.Extent)
}

func (r Rectangle) Accept(visitor Visitor) {
	Polygon{r.Points()}.Accept(visitor)
}

func (r Rectangle) Points() []Point {
	return []Point{r.Origin, P(r.Origin.X, r.Origin.Y+r.Extent.Y),
		r.Origin.Plus(r.Extent), P(r.Origin.X+r.Extent.X, r.Origin.Y)}
}
