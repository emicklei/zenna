package xy

type Image struct {
	URL    string
	Center Point
	Width  int
	Height int
}

func (i Image) Accept(visitor Visitor) {
	visitor.VisitImage(i)
}
