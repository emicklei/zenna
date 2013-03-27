package xy

type Composite struct {
	Shapes []Geometric
}

func (s Composite) Accept(visitor Visitor) {
	visitor.VisitComposite(s)
}
func (s *Composite) Add(shape Geometric) {
	s.Shapes = append(s.Shapes, shape)
}
