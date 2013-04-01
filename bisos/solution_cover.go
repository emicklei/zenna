package main

import (
	. "github.com/emicklei/zenna/xy"
)

func bisos_cover() *Composite {
	group := new(Composite)

	one := NewBisos(1)
	group.Add(TranslateBy{one, BisosDelta.Scaled(-5, 1)})

	two := NewBisos(2)
	group.Add(TranslateBy{two, BisosDelta.Scaled(-3, 1)})

	three := NewBisos(3)
	group.Add(TranslateBy{three, BisosDelta.Scaled(1, 3)})

	four := NewBisos(4)
	group.Add(TranslateBy{four, BisosDelta.Scaled(-5, -1)})

	five := NewBisos(5)
	group.Add(TranslateBy{five, BisosDelta.Scaled(-1, -1)})

	six := NewBisos(6)
	group.Add(TranslateBy{six, BisosDelta.Scaled(4, -2)})

	seven := NewBisos(7)
	group.Add(TranslateBy{seven, BisosDelta.Scaled(-4, -2)})

	eight := NewBisos(8)
	group.Add(TranslateBy{eight, BisosDelta.Scaled(-3, -3)})

	nine := NewBisos(9)
	group.Add(TranslateBy{nine, BisosDelta.Scaled(-1, -3)})

	return group
}
