package main

func AreEquiDistantPoints(p1, p2, p3, p4 Plottable) bool {
	d1 := p1.Distance(p2)
	d2 := p3.Distance(p4)

	return d1 == d2
}
