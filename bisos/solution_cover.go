package main

import (
	. "github.com/emicklei/zenna/xy"
)

func bisos_cover() *Composite {
	group := new(Composite)

	one := NewBisos(1)
	group.Add(TranslateOn{one, BisosDelta.Scaled(-5, 1)})

	two := NewBisos(2)
	group.Add(TranslateOn{two, BisosDelta.Scaled(-3, 1)})

	three := NewBisos(3)
	group.Add(TranslateOn{three, BisosDelta.Scaled(1, 3)})

	four := NewBisos(4)
	group.Add(TranslateOn{four, BisosDelta.Scaled(-5, -1)})

	five := NewBisos(5)
	group.Add(TranslateOn{five, BisosDelta.Scaled(-1, -1)})

	six := NewBisos(6)
	group.Add(TranslateOn{six, BisosDelta.Scaled(4, -2)})

	seven := NewBisos(7)
	group.Add(TranslateOn{seven, BisosDelta.Scaled(-4, -2)})

	eight := NewBisos(8)
	group.Add(TranslateOn{eight, BisosDelta.Scaled(-3, -3)})

	nine := NewBisos(9)
	group.Add(TranslateOn{nine, BisosDelta.Scaled(-1, -3)})

	return group
}
