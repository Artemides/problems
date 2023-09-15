package methods

import (
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p *Point) Scale(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i == 0 {
			continue
		}
		sum += path[i-1].Distance(path[i])
	}
	return sum
}

func (path Path) TranslateBy(offset Point, add bool) {
	var f func(p, q Point) Point
	if add {
		f = Point.Add
	} else {
		f = Point.Sub
	}
	for i := range path {
		path[i] = f(path[i], offset)
	}
}

type ColourPoint struct {
	Point
	Color color.RGBA
}

func (p ColourPoint) Distance(q Point) float64 {
	return p.Point.Distance(q)
}
func (p *ColourPoint) Scale(factor float64) {
	p.Point.Scale(factor)
}
func RunMethods() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	p := ColourPoint{Point{5, 4}, red}
	q := ColourPoint{Point{15, 14}, blue}
	p.Distance(q.Point)

	distance := (*Point).Scale
	distance(&p.Point, 21)
}
