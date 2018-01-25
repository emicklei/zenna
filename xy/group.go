package xy

type Group []Geometric

func (g Group) Accept(visitor Visitor) {
	visitor.VisitGroup(g)
}
func (g *Group) Add(shape Geometric) {
	*g = append(*g, shape)
}
