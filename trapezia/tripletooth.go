package main

import (
	_ "math"
	"os"

	"github.com/ajstarks/svgo"
	. "github.com/emicklei/zenna/svgf"
	. "github.com/emicklei/zenna/xy"
)

func main() {
	canvas := svg.New(os.Stdout)
	canvas.Start(600, 600)
	canvas.Title("TripleTooth")
	canvas.Translate(300, 300)
	canvas.Scale(150)

	painter := NewSVGPainter("", canvas)

	// usePen(painter)
	useTrapezia(painter)

	canvas.Gend() //scale
	canvas.Gend() //translat
	canvas.End()  // canvas
}

func usePen(painter Painter) {
	// red:#FF358B
	// green: #AEEE00
	// blue: #01B0F0

	painter.Style("stroke-width:0.08px;stroke:#01B0F0;fill:#AEEE00")
	p := NewPen(PointZero, P(1, 1))
	p.Move(0, 0.5)
	p.Move(-0.5*Cos30, -0.25)
	p.Move(0, 0.5)
	p.Move(0.5*Cos30, 0.25)
	p.Move(0.5*Cos30, -0.25)
	p.Move(0, -1)
	p.Move(0.5*Cos30, -0.25)
	p.Move(0, 0.5)
	p.Move(0.5*Cos30, -0.25)
	p.Move(0, -0.5)
	p.Move(-0.5*Cos30, -0.25)
	p.Move(-Cos30, 0.5)
	p.Move(-0.5*Cos30, -0.25)
	p.Move(0.5*Cos30, -0.25)
	p.Move(-0.5*Cos30, -0.25)
	p.Move(-0.5*Cos30, 0.25)
	p.Move(0, 0.5)
	p.Move(Cos30, 0.5)
	painter.Paint(Polygon{p.Path()})
}

func useTrapezia(painter Painter) {
	t1 := Trapezium{60, 1}
	dx := t1.Delta().X
	dy := t1.Delta().Y
	// grid
	//painter.Paint(StyleWith{Raster{PointZero, 6, 6, dx, dy}, "stroke-width:0.01px;stroke:gray"})
	// x-y
	//painter.Paint(StyleWith{Raster{PointZero, 6, 6, 0, 0}, "stroke-width:0.01px;stroke:red"})

	painter.Style("stroke-width:0.01px;stroke:blue;fill:blue")

	g0 := new(Composite)
	g0.Add(TranslateOn{P(2*dy/3, 0),
		RotateOn{90, t1}})
	g0.Add(TranslateOn{
		P(-2*dx-(dy/3), 0.5*t1.Length),
		RotateOn{30, t1}})

	g1 := new(Composite)
	g1.Add(g0)
	g1.Add(RotateOn{120, g0})
	g1.Add(RotateOn{240, g0})

	painter.Paint(g1)
	painter.Paint(TranslateOn{P(-t1.Length*Cos30+(dy/3), -2),
		RotateOn{-60, g1}})

}
