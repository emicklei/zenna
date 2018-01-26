package xy

type Text struct {
	Position Point
	Text     string
	Delta    Point
}

func (t Text) Accept(visitor Visitor) {
	visitor.VisitText(t)
}
