package main

import "math"

type Plottable interface {
	GetX() int
	GetY() int
	Distance(Plottable) float64
}

type Point struct {
	x, y int
}

func Initialize(x, y int) Point {
	p1 := Point{x, y}
	return p1
}

func (p *Point) GetX() int {
	return p.x
}

func (p *Point) GetY() int {
	return p.y
}

func (p *Point) Distance(p2 Plottable) float64 {
	return math.Sqrt(math.Pow(float64(p2.GetX()-p.GetX()), 2) + math.Pow(float64(p2.GetY()-p.GetY()), 2))
}
