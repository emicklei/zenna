package main

import (
	"os"

	svg "github.com/ajstarks/svgo"
	"github.com/emicklei/zenna/svgf"
	. "github.com/emicklei/zenna/xy"
)

func main() {
	canvas := svg.New(os.Stdout)
	canvas.Start(2400, 1200)
	canvas.Title("Musik Notez")
	canvas.Translate(200, 200)

	painter := svgf.NewSVGPainter("", canvas)

	painter.Paint(Style{"stroke-width:1px;stroke:red", LineSegment{P(-200, 0), P(200, 0)}})
	painter.Paint(Style{"stroke-width:1px;stroke:red", LineSegment{P(0, 200), P(0, -200)}})

	stick := Style{"stroke-width:3px;stroke:black",
		LineSegment{P(0, 0), P(0, 60)}}
	bubble := Translate{
		P(11, 60),
		Rotate{15,
			Style{"stroke-width:2px;stroke:black;fill:gray",
				Ellipse{PointZero, 12, 8, 0, 0}}}}
	upNote := Group{stick, bubble}

	stick2 := Style{"stroke-width:3px;stroke:black",
		LineSegment{P(0, 0), P(0, -60)}}
	bubble2 := Translate{
		P(-11, -60),
		Rotate{15,
			Style{"stroke-width:2px;stroke:black;fill:gray",
				Ellipse{PointZero, 12, 8, 0, 0}}}}
	downNote := Group{stick2, bubble2}

	/**
	painter.Paint(
		Style{"stroke-width:1px;stroke:gray",
			Raster{PointZero, 2400, 1200, 10, 10}})
	**/

	//painter.Paint(Translate{P(200, -200), upNote})
	//painter.Paint(Translate{P(240, -200), upNote})
	painter.Paint(upNote)
	painter.Paint(downNote)

	{
		bar := Style{"stroke-width:1px;stroke:black",
			LineSegment{P(-60, 51), P(60, 51)}}
		painter.Paint(bar)
	}

	{
		bar := Style{"stroke-width:1px;stroke:black",
			LineSegment{P(-60, 69), P(60, 69)}}
		painter.Paint(bar)
	}

	canvas.Gend() // Translate
	canvas.End()
}
