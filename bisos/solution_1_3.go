package main

import (
	. "github.com/emicklei/zenna/xy"
)

func bisos_1_3() *Composite {
	group := new(Composite)
	style := "stroke-width:10px;stroke:red;fill:none"

	four := NewBisos(style, 4)
	(&four).Mirrored(AxisY(0))
	group.Add(TranslateBy{RotateBy{four, 60}, BisosDelta.Scaled(0, -3)})

	one := NewBisos(style, 1)
	group.Add(TranslateBy{RotateBy{one, 120}, BisosDelta.Scaled(3, 2)})

	eight := NewBisos(style, 8)
	group.Add(TranslateBy{RotateBy{eight, 60}, BisosDelta.Scaled(-2, 1)})

	three := NewBisos(style, 3)
	(&three).Mirrored(AxisY(0))
	group.Add(TranslateBy{RotateBy{three, -60}, BisosDelta.Scaled(-2, 1)})

	return group
}
