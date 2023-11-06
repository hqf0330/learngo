package main

type Point2 struct {
	X, y int
}

type Circle struct {
	Center Point2
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

func main() {
	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.y = 8
	w.Circle.Radius = 5
	w.Spokes = 20

}
