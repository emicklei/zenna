package main

import (
	. "github.com/emicklei/zenna/xy"
)

func bisos_1_3() *Composite {
	group := new(Composite)

	four := NewBisos(4)
	(&four).Mirrored(AxisY(0))
	group.Add(Translate{BisosDelta.Scaled(0, -3), Rotate{60, four}})

	one := NewBisos(1)
	group.Add(Translate{BisosDelta.Scaled(3, 2), Rotate{120, one}})

	eight := NewBisos(8)
	group.Add(Translate{BisosDelta.Scaled(-2, 1), Rotate{60, eight}})

	three := NewBisos(3)
	(&three).Mirrored(AxisY(0))
	group.Add(Translate{BisosDelta.Scaled(-2, 1), Rotate{-60, three}})

	return group
}
