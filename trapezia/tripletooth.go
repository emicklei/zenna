package main

import (
	"github.com/ajstarks/svgo"
	. "github.com/emicklei/zenna/svgf"
	. "github.com/emicklei/zenna/xy"
	_ "math"
	"os"
)

func main() {
	canvas := svg.New(os.Stdout)
	canvas.Start(600, 600)
	canvas.Title("Trapezium")
	canvas.Translate(300, 300)
	canvas.Scale(100)

	painter := NewSVGPainter("", canvas)
	t1 := Trapezium{60, 1}
	dx := t1.Delta().X
	dy := t1.Delta().Y
	// grid
	//painter.Paint(StyleWith{Raster{PointZero, 6, 6, dx, dy}, "stroke-width:0.01px;stroke:gray"})
	// x-y
	//painter.Paint(StyleWith{Raster{PointZero, 6, 6, 0, 0}, "stroke-width:0.01px;stroke:red"})

	painter.Style("stroke-width:0.01px;stroke:black;fill:none")

	g0 := new(Composite)
	g0.Add(TranslateBy{RotateBy{t1, 90}, P(2*dy/3, 0)})
	g0.Add(TranslateBy{RotateBy{t1, 30}, P(-2*dx-(dy/3), 0.5*t1.Length)})
	painter.Paint(g0)

	painter.Paint(RotateBy{g0, 120})
	painter.Paint(RotateBy{g0, 240})

	canvas.Gend() //scale
	canvas.Gend() //translat
	canvas.End()  // canvas
}
