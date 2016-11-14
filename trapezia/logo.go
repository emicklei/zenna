package main

import (
	_ "math"
	"os"

	"github.com/ajstarks/svgo"
	. "github.com/emicklei/zenna/svgf"
	. "github.com/emicklei/zenna/xy"
)

// http://colorschemedesigner.com/

func main() {
	canvas := svg.New(os.Stdout)
	canvas.Start(600, 600)
	canvas.Title("Trapezium")
	canvas.Translate(300, 300)
	canvas.Scale(100)

	painter := NewSVGPainter("", canvas)

	t1 := Trapezium{60, 1}

	// grid
	painter.Paint(Style{"stroke-width:0.01px;stroke:gray",
		Raster{PointZero, 6, 6, t1.Delta().X, t1.Delta().Y}})
	// x-y
	painter.Paint(Style{"stroke-width:0.01px;stroke:red",
		Raster{PointZero, 6, 6, 0, 0}})

	g0 := new(Composite)
	// green
	g0.Add(Style{"fill:#9FEE00;stroke:#9FEE00;stroke-width:0.01px", t1})

	g0.Add(Translate{t1.Delta().Multiply(2), Rotate{-120, t1}})
	/**
	g0.Add(
		Style{"fill:#1240AB;stroke:#1240AB;stroke-width:0.01px",
			Translate{t1.Delta().Scaled(0, 0),
				Rotate{120, t1}}})
	**/

	//t3 := Style{"fill:#FFAA00;stroke:#FFAA00;stroke-width:0.01px", Trapezium{60, 1}}
	//g0.Add(Translate{Point{0.75, -HCos30}, Rotate{120, t3}})

	painter.Paint(g0)

	g1 := new(Composite)
	g1.Add(g0)

	g1.Add(Translate{Point{0.5, -Cos30}, Rotate{60, g0}})
	g1.Add(Translate{Point{0.5, -Cos30}, Rotate{120, g0}})
	g1.Add(Translate{Point{2, 0}, Rotate{180, g0}})
	g1.Add(Translate{Point{1.5, Cos30}, Rotate{240, g0}})
	g1.Add(Translate{Point{0.5, Cos30}, Rotate{300, g0}})

	//g1.Accept(painter)

	/**
		TranslateBy{g1, Point{Cos30 + 1.5, Cos30 * 3 / 2}}.Accept(painter)

		TranslateBy{g1, Point{-Cos30 - 1.5, Cos30 * 3 / 2}}.Accept(painter)

		TranslateBy{g1, Point{Cos30 + 1.5, Cos30 * 9 / 1.95}}.Accept(painter)

		TranslateBy{g1, Point{-Cos30 - 1.5, Cos30 * 9 / 1.95}}.Accept(painter)

		TranslateBy{g1, Point{0, Cos30 * 12 / 1.95}}.Accept(painter)
	**/
	canvas.Gend()
	canvas.Gend()
	canvas.End()
}
