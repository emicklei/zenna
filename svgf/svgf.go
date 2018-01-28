package svgf

import (
	"fmt"
	"strings"

	svg "github.com/ajstarks/svgo"
)

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

// Image places at x,y (upper left hand corner), the image with
// width w, and height h, referenced at link, with optional style.
// Standard Reference: http://www.w3.org/TR/SVG11/struct.html#ImageElement
func (svgf *SVGF) Image(xc, yc float64, w, h float64, href, s string) {
	// svg.printf(`<image %s %s %s`, dim(x, y, w, h), href(link), endstyle(s, emptyclose))
	fmt.Fprintf(svgf.Writer, `<image x="%fpx" y="%fpx" width="%f" height="%f" href="%s" %s/>`, xc+(float64(w)/2), yc-(float64(h)/2), w, h, href, style(s))
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
