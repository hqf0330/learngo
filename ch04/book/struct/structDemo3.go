package main

import "fmt"

type Point3 struct {
	X, Y int
}

type Circle3 struct {
	Point3
	Radius int
}

type Wheel3 struct {
	Circle3
	Spokes int
}

func main() {
	var w Wheel3
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20

	w2 := Wheel3{
		Circle3: Circle3{
			Point3: Point3{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}

	fmt.Printf("%#v\n", w2)
	w2.X = 42
	fmt.Printf("%#v\n", w2)

}
