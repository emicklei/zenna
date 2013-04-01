package main

import (
	"github.com/ajstarks/svgo"
	. "github.com/emicklei/zenna/svgf"
	. "github.com/emicklei/zenna/xy"
	"os"
	_ "math"
)

// http://colorschemedesigner.com/

func main() {
	canvas := svg.New(os.Stdout)
	canvas.Start(600, 600)
	canvas.Title("Trapezium")
	canvas.Translate(300, 300)
	canvas.Scale(100)

	painter := NewSVGPainter("", canvas)
	t1 := Trapezium{60,1}
	// grid
	painter.Paint(StyleWith{Grid{PointZero, 6, 6, t1.Delta().X , t1.Delta().Y }, "stroke-width:0.01px;stroke:gray"})
	// x-y
	painter.Paint(StyleWith{Grid{PointZero, 6, 6, 0 , 0 }, "stroke-width:0.01px;stroke:red"})
	
    g0 := new(Composite)
	g0.Add(StyleWith{t1,"fill:9FEE00;stroke:9FEE00;stroke-width:0.01px"})
	
	g0.Add(StyleWith{TranslateBy{RotateBy{t1, 120}, t1.Delta().Scaled(0, 0)},"fill:1240AB;stroke:1240AB;stroke-width:0.01px"})


	//t3 := Trapezium{"fill:FFAA00;stroke:FFAA00;stroke-width:0.01px", 60, 1}
	//g0.Add(TranslateBy{RotateBy{t3, -120}, Point{0.75, -HCos30}})

	painter.Paint(g0)

	//g1 := new(Composite)
	//g1.Add(g0)

	//g1.Add(TranslateBy{RotateBy{g0, 60}, Point{0.5, -Cos30}})

	//g1.Add(TranslateBy{RotateBy{g0, 120}, Point{1.5, -Cos30}})

	//g1.Add(TranslateBy{RotateBy{g0, 180}, Point{2, 0}})

	//g1.Add(TranslateBy{RotateBy{g0, 240}, Point{1.5, Cos30}})

	//g1.Add(TranslateBy{RotateBy{g0, 300}, Point{0.5, Cos30}})

	//g1.Accept(painter)

	//TranslateBy{g1, Point{Cos30 + 1.5, Cos30 * 3 / 2}}.Accept(painter)

	//TranslateBy{g1, Point{-Cos30 - 1.5, Cos30 * 3 / 2}}.Accept(painter)

	//TranslateBy{g1, Point{Cos30 + 1.5, Cos30 * 9 / 1.95}}.Accept(painter)

	//TranslateBy{g1, Point{-Cos30 - 1.5, Cos30 * 9 / 1.95}}.Accept(painter)

	//TranslateBy{g1, Point{0, Cos30 * 12 / 1.95}}.Accept(painter)

	canvas.Gend()
	canvas.Gend()
	canvas.End()
}
