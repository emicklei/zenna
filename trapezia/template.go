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
	// grid
	painter.Paint(StyleWith{Raster{PointZero, 6, 6, t1.Delta().X, t1.Delta().Y}, "stroke-width:0.01px;stroke:gray"})
	// x-y
	painter.Paint(StyleWith{Raster{PointZero, 6, 6, 0, 0}, "stroke-width:0.01px;stroke:red"})

	canvas.Gend() //scale
	canvas.Gend() //translat
	canvas.End()  // canvas
}
