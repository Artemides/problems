package structs

type Circle struct {
	X, Y, Radius int
}
type Wheel struct {
	X, Y, Radius, Spokes int
}

type Point struct {
	X, Y int
}
type Circle2 struct {
	Center Point
	Radius int
}
type Circle3 struct {
	Point
	Radius int
}

type Wheel2 struct {
	Circle Circle2
	Spokes int
}

type Wheel3 struct {
	Circle3
	Spokes int
}

func Run1() {
	var wheel Wheel
	wheel.X = 10
	wheel.Y = 20
	wheel.Radius = 40
	wheel.Spokes = 20
}

func Run2() {
	var wheel Wheel2
	wheel.Circle.Center.X = 10
	wheel.Circle.Center.Y = 20
	wheel.Circle.Radius = 39
	wheel.Spokes = 12
}

func Run3() {
	var wheel Wheel3
	wheel.X = 10
	wheel.Y = 30
	wheel.Radius = 25
	wheel.Spokes = 39
}
