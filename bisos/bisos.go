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

	vis := NewSVGPainter("", canvas)

	switch os.Args[1] {
	case "cover":
		ScaleBy{bisos_cover(), 0.5}.Accept(vis)
	case "1_3":
		ScaleBy{bisos_1_3(), 0.5}.Accept(vis)
	}

	canvas.Gend()
	canvas.End()
}
