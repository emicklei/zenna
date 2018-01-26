package svgf

import (
	"fmt"
	"strings"

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

// Code below exists until SVG package has support for Float64 numbers

type SVGF struct {
	*svg.SVG
}

// Circle centered at x,y, with radius r, with optional style.
// Standard Reference: http://www.w3.org/TR/SVG11/shapes.html#CircleElement
func (svgf *SVGF) Circle(x float64, y float64, r float64, s ...string) {
	fmt.Fprintf(svgf.Writer, `<circle cx="%v" cy="%v" r="%v" %s`, x, y, r, endstyle(s, emptyclose))
}

// Line draws a straight line between two points, with optional style.
// Standard Reference: http://www.w3.org/TR/SVG11/shapes.html#LineElement
func (svgf *SVGF) Line(x1 float64, y1 float64, x2 float64, y2 float64, s ...string) {
	fmt.Fprintf(svgf.Writer, `<line x1="%v" y1="%v" x2="%v" y2="%v" %s`, x1, y1, x2, y2, endstyle(s, emptyclose))
}

// Polygon draws a series of line segments using an array of x, y coordinates, with optional style.
// Standard Reference: http://www.w3.org/TR/SVG11/shapes.html#PolygonElement
func (svgf *SVGF) Polygon(x []float64, y []float64, s ...string) {
	fmt.Fprintf(svgf.Writer, `<polygon style="%s" points="`, s[0])
	for i := 0; i < len(x); i++ {
		fmt.Fprintf(svgf.Writer, "%v,%v ", x[i], y[i])
	}
	fmt.Fprintf(svgf.Writer, `" />`)
}

// https://www.w3.org/TR/SVG11/shapes.html#EllipseElement
func (svgf *SVGF) Ellipse(x1 float64, y1 float64, rx float64, ry float64, s ...string) {
	fmt.Fprintf(svgf.Writer, `<ellipse cx="%f" cy="%f" rx="%f" ry="%f" %s`, x1, -y1, rx, ry, endstyle(s, emptyclose))
}

// https://www.w3.org/TR/SVG11/shapes.html#EllipseElement
func (svgf *SVGF) Text(x float64, y float64, dx float64, dy float64, t string, s string) {
	fmt.Fprintf(svgf.Writer, `<text x="%f" y="%f" dx="%f" dy="%f" %s>%s</text>`, x, -y, dx, -dy, style(s), t)
}

// ----- Copied Code because not Exported

const emptyclose = "/>\n"

// endstyle modifies an SVG object, with either a series of name="value" pairs,
// or a single string containing a style
func endstyle(s []string, endtag string) string {
	if len(s) > 0 {
		nv := ""
		for i := 0; i < len(s); i++ {
			if strings.Index(s[i], "=") > 0 {
				nv += (s[i]) + " "
			} else {
				nv += style(s[i])
			}
		}
		return nv + endtag
	}
	return endtag

}

// style returns a style name,attribute string
func style(s string) string {
	if len(s) > 0 {
		return fmt.Sprintf(`style="%s"`, s)
	}
	return s
}
