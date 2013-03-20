package svgf

import (
	"fmt"
	"github.com/ajstarks/svgo"
	. "github.com/emicklei/zenna/xy"
	"strings"
)

type Painter struct {
	style  string // current style
	canvas *svg.SVG
}

func NewSVGPainter(style string, canvas *svg.SVG) Painter {
	return Painter{style,canvas}
}

func (p *Painter) Style(newStyle string) {
	p.style = newStyle
}

func (p Painter) VisitTranslateBy(t TranslateBy) {
	p.canvas.Gtransform(fmt.Sprintf(`translate(%f,%f)`, t.Translation.X, t.Translation.Y))
	t.Shape.Accept(p)
	p.canvas.Gend()
}
func (p Painter) VisitScaleBy(s ScaleBy) {
	p.canvas.Scale(s.Scale)
	s.Shape.Accept(p)
	p.canvas.Gend()
}
func (p Painter) VisitRotateBy(r RotateBy) {
	p.canvas.Rotate(r.Angle)
	r.Shape.Accept(p)
	p.canvas.Gend()
}
func (p Painter) VisitShapeComposite(s ShapeComposite) {
	for _, each := range s.Shapes {
		each.Accept(p)
	}
}
func (p Painter) VisitPolygon(poly Polygon) {
	// Hack: write directly to Stdout
	fmt.Printf(`<polygon style="%s" points="`, p.style)
	for _, each := range poly.Points {
		fmt.Printf("%f,%f ", each.X, each.Y)
	}
	fmt.Println(`" />`)
}
func (p Painter) VisitStyleWith(s StyleWith) {
	old := p.style
	(&p).style = s.Style
	s.Shape.Accept(p)
	(&p).style = old // use defer?
}

// Code below exists until SVG package has support for Float64 numbers

type SVGF struct {
	*svg.SVG
}

// Line draws a straight line between two points, with optional style.
// Standard Reference: http://www.w3.org/TR/SVG11/shapes.html#LineElement
func (svgf *SVGF) Line(x1 float64, y1 float64, x2 float64, y2 float64, s ...string) {
	fmt.Fprintf(svgf.Writer, `<line x1="%f" y1="%f" x2="%f" y2="%f" %s`, x1, y1, x2, y2, endstyle(s, emptyclose))
}

// Polygon draws a series of line segments using an array of x, y coordinates, with optional style.
// Standard Reference: http://www.w3.org/TR/SVG11/shapes.html#PolygonElement
func (svg *SVGF) Polygon(x []float64, y []float64, s ...string) {
	fmt.Fprintf(svg.Writer, `<polygon style="%s" points="`, s[0])
	for i := 0; i < len(x); i++ {
		fmt.Fprintf(svg.Writer, "%f,%f ", x[i], y[i])
	}
	fmt.Fprintf(svg.Writer, `" />`)
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
