package xy

type Curve struct {
	block func(param float64) Point
}

func (c Curve) PointAt(param float64) Point {
	return c.block(param)
}
