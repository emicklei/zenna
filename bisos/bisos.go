package main

import (
	"github.com/ajstarks/svgo"
	. "github.com/emicklei/zenna/svgf"
	. "github.com/emicklei/zenna/xy"
	"os"
)

func main() {
	if len(os.Args) != 2 || os.Args[0] == "$1" {
		println("bisos <solution>")
		return
	}

	canvas := svg.New(os.Stdout)
	canvas.Start(800, 600)
	canvas.Title(os.Args[1])
	canvas.Translate(400, 300)

	painter := NewSVGPainter("", canvas)

	// grid
	//	painter.Paint(StyleWith{Grid{PointZero, 800, 600, BisosDelta.X/2, BisosDelta.Y/2}, "stroke-width:1px;stroke:gray"})
	// x-y
	//	painter.Paint(StyleWith{Grid{PointZero, 800, 600, 0 , 0 }, "stroke-width:1px;stroke:red"})

	switch os.Args[1] {
	case "cover":
		painter.Paint(StyleWith{ScaleBy{bisos_cover(), 0.5}, "stroke:blue;fill:none"})
	case "1_3":
		painter.Paint(StyleWith{ScaleBy{bisos_1_3(), 0.5}, "stroke:red;fill:none"})
		//painter.Paint(StyleWith{"stroke:red;fill:none", ScaleBy{0.5, bisos_1_3()}})
	}

	canvas.Gend()
	canvas.End()
}
