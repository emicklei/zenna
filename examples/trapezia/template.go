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
	canvas.Title("Trapezium")
	canvas.Translate(300, 300)
	canvas.Scale(100)

	painter := NewSVGPainter("", canvas)
	t1 := Trapezium{60, 1}
	// grid
	painter.Paint(Style{"stroke-width:0.01px;stroke:gray", Raster{PointZero, 6, 6, t1.Delta().X, t1.Delta().Y}})
	// x-y
	painter.Paint(Style{"stroke-width:0.01px;stroke:red", Raster{PointZero, 6, 6, 0, 0}})

	canvas.Gend() //scale
	canvas.Gend() //translat
	canvas.End()  // canvas
}
