package xy

type ShapeComposite struct {
	Shapes []Geometric2D
}

func (s ShapeComposite) Accept(visitor Visitor2D) {
	visitor.VisitShapeComposite(s)
}
func (s *ShapeComposite) Add(shape Geometric2D) {
	s.Shapes = append(s.Shapes, shape)
}
