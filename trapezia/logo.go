package main

import (
	"github.com/ajstarks/svgo"
	. "github.com/emicklei/zenna/svgf"
	. "github.com/emicklei/zenna/xy"
	"os"
)

// http://colorschemedesigner.com/

func main() {
	canvas := svg.New(os.Stdout)
	canvas.Start(1200, 1200)
	canvas.Title("Trapezium")
	canvas.Translate(600, 400)
	canvas.Scale(100)

	vis := NewSVGPainter("", canvas)
	//canvas.Script("application/javascript", `function timedRefresh(timeoutPeriod) {
	//setTimeout("location.reload(true);",timeoutPeriod);
	//}`)

	//Grid{Point{0, 0}, 10, 10, 1, 1, "stroke:gray;stroke-width:0.01px"}.DrawOn(canvas)

	g0 := new(Composite)

	t1 := Trapezium{"fill:9FEE00;stroke:9FEE00;stroke-width:0.01px", 60, 1}
	g0.Add(t1)

	t2 := Trapezium{"fill:1240AB;stroke:1240AB;stroke-width:0.01px", 60, 1}
	g0.Add(TranslateBy{RotateBy{t2, 120}, t2.Delta().Scaled(0, 2)})

	t3 := Trapezium{"fill:FFAA00;stroke:FFAA00;stroke-width:0.01px", 60, 1}
	g0.Add(TranslateBy{RotateBy{t3, -120}, Point{0.75, -HCos30}})

	g1 := new(Composite)
	g1.Add(g0)

	g1.Add(TranslateBy{RotateBy{g0, 60}, Point{0.5, -Cos30}})

	g1.Add(TranslateBy{RotateBy{g0, 120}, Point{1.5, -Cos30}})

	g1.Add(TranslateBy{RotateBy{g0, 180}, Point{2, 0}})

	g1.Add(TranslateBy{RotateBy{g0, 240}, Point{1.5, Cos30}})

	g1.Add(TranslateBy{RotateBy{g0, 300}, Point{0.5, Cos30}})

	g1.Accept(vis)

	TranslateBy{g1, Point{Cos30 + 1.5, Cos30 * 3 / 2}}.Accept(vis)

	TranslateBy{g1, Point{-Cos30 - 1.5, Cos30 * 3 / 2}}.Accept(vis)

	TranslateBy{g1, Point{Cos30 + 1.5, Cos30 * 9 / 1.95}}.Accept(vis)

	TranslateBy{g1, Point{-Cos30 - 1.5, Cos30 * 9 / 1.95}}.Accept(vis)

	TranslateBy{g1, Point{0, Cos30 * 12 / 1.95}}.Accept(vis)

	canvas.Gend()
	canvas.Gend()
	canvas.End()
}
