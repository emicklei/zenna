package xy

type RoundedRectangle struct {
	Rectangle
	XRadius, YRadius float64
}

func (r RoundedRectangle) Accept(visitor Visitor) {
	visitor.VisitRoundedRectangle(r)
}
