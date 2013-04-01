package main

import (
	. "github.com/emicklei/zenna/xy"
)

var BisosDelta = Point{1, Tan60}.Multiply(float64(100))

type Bisos struct {
	points []Point
}

func NewBisos(id byte) Bisos {
	p := NewPen(Point{}, BisosDelta)
	switch id {
	case 1:
		bisos_1(&p)
	case 2:
		bisos_2(&p)
	case 3:
		bisos_3(&p)
	case 4:
		bisos_4(&p)
	case 5:
		bisos_5(&p)
	case 6:
		bisos_6(&p)
	case 7:
		bisos_7(&p)
	case 8:
		bisos_8(&p)
	case 9:
		bisos_9(&p)
	}
	return Bisos{p.Path()}
}

func bisos_1(p *Pen) {
	p.Move(2, 2)
	p.Move(4, 0)
	p.Move(-1, -1)
	p.Move(-2, 0)
	p.Move(-1, -1)
	p.Move(-2, 0)
}

func bisos_2(p *Pen) {
	p.Move(1, 1)
	p.Move(2, 0)
	p.Move(1, 1)
	p.Move(2, -2)
	p.Move(-6, 0)
}

func bisos_3(p *Pen) {
	p.Move(2, 0)
	p.Move(3, -3)
	p.Move(-2, 0)
	p.Move(-3, 3)
}

func bisos_4(p *Pen) {
	p.Move(-1, 1)
	p.Move(1, 1)
	p.Move(4, 0)
	p.Move(1, -1)
	p.Move(-4, 0)
	p.Move(-1, -1)
}

func bisos_5(p *Pen) {
	p.Move(-1, 1)
	p.Move(2, 0)
	p.Move(-1, 1)
	p.Move(4, 0)
	p.Move(-2, -2)
	p.Move(-2, 0)
}

func bisos_6(p *Pen) {
	p.Move(-1, 1)
	p.Move(-2, 0)
	p.Move(2, 2)
	p.Move(1, -1)
	p.Move(2, 0)
	p.Move(-2, -2)
}

func bisos_7(p *Pen) {
	p.Move(-1, 1)
	p.Move(1, 1)
	p.Move(2, 0)
	p.Move(1, -1)
	p.Move(-1, -1)
	p.Move(-2, 0)
}

func bisos_8(p *Pen) {
	p.Move(-1, 1)
	p.Move(2, 0)
	p.Move(1, 1)
	p.Move(2, 0)
	p.Move(1, -1)
	p.Move(-2, 0)
	p.Move(-1, -1)
	p.Move(-2, 0)
}

func bisos_9(p *Pen) {
	p.Move(1, 1)
	p.Move(2, 0)
	p.Move(-1, 1)
	p.Move(2, 0)
	p.Move(1, -1)
	p.Move(-1, -1)
	p.Move(-4, 0)
}

func (b *Bisos) Mirrored(axis Axis) {
	b.points = Mirrored(b.points, axis)
}

func (b Bisos) Accept(visitor Visitor) {
	Polygon{b.points}.Accept(visitor)
}
