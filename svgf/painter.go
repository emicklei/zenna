package svgf

import (
	"fmt"

	"github.com/ajstarks/svgo"
	. "github.com/emicklei/zenna/xy"
)

// Painter is a Geometric Visitor that renders SVG documents
// It translates the Zenna coordinate system (Y=up) to that of SVG (Y=down)
type Painter struct {
	style  string // current style
	canvas *svg.SVG
}

func NewSVGPainter(style string, canvas *svg.SVG) Painter {
	return Painter{style, canvas}
}

func (p Painter) Paint(g Geometric) {
	g.Accept(p)
}

func (p *Painter) Style(newStyle string) {
	p.style = newStyle
}

func (p Painter) VisitCircle(c Circle) {
	(&SVGF{p.canvas}).Circle(c.Center.X, -c.Center.Y, c.Radius, p.style)
}

func (p Painter) VisitTranslate(t Translate) {
	p.canvas.Gtransform(fmt.Sprintf(`translate(%f,%f)`, t.Translation.X, -t.Translation.Y))
	t.Shape.Accept(p)
	p.canvas.Gend()
}
func (p Painter) VisitScale(s Scale) {
	p.canvas.Scale(s.Scale)
	s.Shape.Accept(p)
	p.canvas.Gend()
}
func (p Painter) VisitRotate(r Rotate) {
	p.canvas.Rotate(-r.Angle)
	r.Shape.Accept(p)
	p.canvas.Gend()
}
func (p Painter) VisitGroup(s Group) {
	for _, each := range s {
		each.Accept(p)
	}
}
func (p Painter) VisitPolygon(poly Polygon) {
	fmt.Fprintf(p.canvas.Writer, `<polygon style="%s" points="`, p.style)
	for _, each := range poly.Points {
		fmt.Fprintf(p.canvas.Writer, "%v,%v ", each.X, -each.Y)
	}
	fmt.Fprintln(p.canvas.Writer, `" />`)
}
func (p Painter) VisitStyle(s Style) {
	old := p.style
	(&p).style = s.Style
	s.Shape.Accept(p)
	(&p).style = old // use defer?
}

func (p Painter) VisitLineSegment(l LineSegment) {
	(&SVGF{p.canvas}).Line(l.Begin.X, -l.Begin.Y, l.End.X, -l.End.Y, p.style)
}

// VisitEllipse is part of Visitor
func (p Painter) VisitEllipse(e Ellipse) {
	(&SVGF{p.canvas}).Ellipse(e.Center.X, e.Center.Y, e.RadiusX, e.RadiusY, p.style)
}

func (p Painter) VisitRectangle(r Rectangle) {
	// todo
}

func (p Painter) VisitRoundedRectangle(r RoundedRectangle) {
	// todo
}

func (p Painter) VisitText(t Text) {
	(&SVGF{p.canvas}).Text(t.Position.X, t.Position.Y, t.Delta.X, t.Delta.Y, t.Text, p.style)
}

func (p Painter) VisitImage(i Image) {
	(&SVGF{p.canvas}).Image(i.Center.X, i.Center.Y, float64(i.Width), float64(i.Height), i.URL, p.style)
}
