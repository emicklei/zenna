package main

import (
	"os"
	"strconv"

	svg "github.com/ajstarks/svgo"
	"github.com/emicklei/zenna/svgf"
	. "github.com/emicklei/zenna/xy"
)

const (
	stickStyle   = "stroke-width:2px;stroke:black;stroke-linecap:round"
	bubbleStyle  = "stroke-width:2px;stroke:black;fill:black"
	barStyle     = "stroke-width:2px;stroke:black"
	barWidth     = 800
	auxBarStyle  = "stroke-width:3px;stroke:black"
	stickSpaceX  = 36.0
	bubbleSpaceY = 17.0
	c            = 0
)

var upBar = Style{auxBarStyle,
	LineSegment{P(-6, 60), P(29, 60)}}
var downBar = Style{auxBarStyle,
	LineSegment{P(-6-23, 60), P(29-23, 60)}}

func main() {
	canvas := svg.New(os.Stdout)
	canvas.Start(2400, 1200)
	canvas.Title("Musik Notez")
	canvas.Translate(barWidth/2, 200)

	painter := svgf.NewSVGPainter("", canvas)

	painter.Paint(Style{"stroke-width:1px;stroke:red", LineSegment{P(-200, 0), P(200, 0)}})
	painter.Paint(Style{"stroke-width:1px;stroke:red", LineSegment{P(0, 200), P(0, -200)}})

	stick := Style{stickStyle,
		LineSegment{P(0, 0), P(0, 60)}}
	bubble := Translate{
		P(11.5, 60),
		Rotate{15,
			Style{bubbleStyle,
				Ellipse{PointZero, 12, 8, 0, 0}}}}
	upNote := Group{stick, bubble}

	stick2 := Style{stickStyle,
		LineSegment{P(0, 60), P(0, 120)}}
	bubble2 := Translate{
		P(-11.5, 60),
		Rotate{15,
			Style{bubbleStyle,
				Ellipse{PointZero, 12, 8, 0, 0}}}}
	downNote := Group{stick2, bubble2}

	bars(painter)

	for t := 1; t < 10; t++ {
		note := c + t - 1
		off := P(float64(t)*stickSpaceX-(stickSpaceX/2), float64(t)*bubbleSpaceY*0.5)
		auxbars(painter, off, note)
		painter.Paint(Translate{off, upNote})
		painter.Paint(Translate{off, Text{Text: strconv.Itoa(note)}})
	}
	for t := 0; t < 10; t++ {
		note := c - t
		off := P(float64(t)*-stickSpaceX, float64(t)*bubbleSpaceY*-0.5)
		auxbars(painter, off, note)
		painter.Paint(Translate{off, downNote})
		painter.Paint(Translate{off, Text{Text: strconv.Itoa(note)}})
	}

	canvas.Gend() // Translate
	canvas.End()
}

func auxbars(painter svgf.Painter, offset Point, note int) {
	for t := 0; t < 10; t++ {
		delta := P(0, float64(t)*bubbleSpaceY*0.5)
		painter.Paint(Translate{offset.Plus(delta), upBar})
	}
}

func bars(painter svgf.Painter) {
	for h := 0; h < 5; h++ {
		bar := Style{barStyle,
			LineSegment{P(-barWidth/2, float64(h)*bubbleSpaceY+34.5), P(barWidth/2, float64(h)*bubbleSpaceY+34.5)}}
		painter.Paint(bar)
	}
}
