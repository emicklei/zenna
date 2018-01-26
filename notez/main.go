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
	barWidth     = 1200
	auxBarStyle  = "stroke-width:2px;stroke:black"
	noteStyle    = "font-weight:bold"
	stickSpaceX  = 42.0
	bubbleSpaceY = 17.0
	c            = 0
	downTextY    = -50
	downTextDY   = 16
	downTextDX   = -18
	upTextY      = 160
	upTextDY     = 100
	upTextDX     = 8
)

var upBar = Style{auxBarStyle,
	LineSegment{P(-6, 60), P(29, 60)}}
var downBar = Style{auxBarStyle,
	LineSegment{P(-6-23, 60), P(29-23, 60)}}

func main() {
	canvas := svg.New(os.Stdout)
	canvas.Start(2400, 1200)
	canvas.Title("Musik Notez")
	painter := svgf.NewSVGPainter("", canvas)

	canvas.Translate(barWidth/2, 240)
	key(painter, GnoteName)
	canvas.Gend() // Translate

	canvas.Translate(barWidth/2, 480)
	key(painter, FnoteName)
	canvas.Gend() // Translate

	canvas.End()
}

func key(painter svgf.Painter, nameFunc func(int) string) {
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

	for t := 1; t < 14; t++ {
		note := c + t
		off := P(float64(t)*stickSpaceX-(stickSpaceX/2), float64(t)*bubbleSpaceY*0.5)
		auxbars(painter, off, note)
		painter.Paint(Translate{off, upNote})
		painter.Paint(Translate{P(off.X+upTextDX, off.Y+upTextDY), Style{noteStyle, Text{Text: nameFunc(note)}}})
	}
	for t := 0; t < 14; t++ {
		note := c - t
		off := P(float64(t)*-stickSpaceX, float64(t)*bubbleSpaceY*-0.5)
		auxbars(painter, off, note)
		painter.Paint(Translate{off, downNote})
		painter.Paint(Translate{P(off.X+downTextDX, off.Y+downTextDY), Style{noteStyle, Text{Text: nameFunc(note)}}})
	}
}

func auxbars(painter svgf.Painter, offset Point, note int) {
	if note < -4 {
		for t := -2; t > note+2; t -= 2 {
			painter.Paint(Translate{P(offset.X, float64(t-3)*bubbleSpaceY*0.5), downBar})
		}
	}
	if note > 6 {
		for t := 4; t < note-2; t += 2 {
			painter.Paint(Translate{P(offset.X, float64(t+3)*bubbleSpaceY*0.5), upBar})
		}
	}
}

func bars(painter svgf.Painter) {
	for h := 0; h < 5; h++ {
		bar := Style{barStyle,
			LineSegment{P(-barWidth/2, float64(h)*bubbleSpaceY+34.5), P(barWidth/2, float64(h)*bubbleSpaceY+34.5)}}
		painter.Paint(bar)
	}
}

func GnoteName(note int) string {
	n := note % 7
	if n < 0 {
		n = 7 - (-note % 7)
	}
	switch n {
	case 0:
		return "A"
	case 1:
		return "B"
	case 2:
		return "C"
	case 3:
		return "D"
	case 4:
		return "E"
	case 5:
		return "F"
	case 6:
		return "G"
	}
	return strconv.Itoa(note)
}

func FnoteName(note int) string {
	n := note % 7
	if n < 0 {
		n = 7 - (-note % 7)
	}
	switch n {
	case 0:
		return "C"
	case 1:
		return "D"
	case 2:
		return "E"
	case 3:
		return "F"
	case 4:
		return "G"
	case 5:
		return "A"
	case 6:
		return "B"
	}
	return strconv.Itoa(note)
}
