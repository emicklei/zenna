package main

import (
	. "github.com/emicklei/zenna/xy"
)

func bisos_1_3() *Composite {
	group := new(Composite)

	four := NewBisos(4)
	(&four).Mirrored(AxisY(0))
	group.Add(TranslateOn{RotateBy{four, 60}, BisosDelta.Scaled(0, -3)})

	one := NewBisos(1)
	group.Add(TranslateOn{RotateBy{one, 120}, BisosDelta.Scaled(3, 2)})

	eight := NewBisos(8)
	group.Add(TranslateOn{RotateBy{eight, 60}, BisosDelta.Scaled(-2, 1)})

	three := NewBisos(3)
	(&three).Mirrored(AxisY(0))
	group.Add(TranslateOn{RotateBy{three, -60}, BisosDelta.Scaled(-2, 1)})

	return group
}
