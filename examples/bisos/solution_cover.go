package main

import (
	. "github.com/emicklei/zenna/xy"
)

func bisos_cover() *Group {
	group := new(Group)

	one := NewBisos(1)
	group.Add(Translate{BisosDelta.Scaled(-5, 1), one})

	two := NewBisos(2)
	group.Add(Translate{BisosDelta.Scaled(-3, 1), two})

	three := NewBisos(3)
	group.Add(Translate{BisosDelta.Scaled(1, 3), three})

	four := NewBisos(4)
	group.Add(Translate{BisosDelta.Scaled(-5, -1), four})

	five := NewBisos(5)
	group.Add(Translate{BisosDelta.Scaled(-1, -1), five})

	six := NewBisos(6)
	group.Add(Translate{BisosDelta.Scaled(4, -2), six})

	seven := NewBisos(7)
	group.Add(Translate{BisosDelta.Scaled(-4, -2), seven})

	eight := NewBisos(8)
	group.Add(Translate{BisosDelta.Scaled(-3, -3), eight})

	nine := NewBisos(9)
	group.Add(Translate{BisosDelta.Scaled(-1, -3), nine})

	return group
}
