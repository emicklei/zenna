package xy

func (p Point) Extent(e Point) Rectangle {
	return Rectangle{p, e}
}
