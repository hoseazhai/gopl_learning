package main

import (
	"fmt"
	"four_struct/trees"
)

func main() {
	println("ttt")
	a := []int{1, 3, 2, 8}
	fmt.Println(trees.Sort(a))

	type Point struct {
		X, Y int
	}
	type Circle struct {
		Point
		Radius int
	}

	type Wheel struct {
		Circle
		Spokes int
	}

	var w Wheel
	w = Wheel{Circle{Point{8, 8}, 5}, 20}
	w = Wheel{
		Circle: Circle{
			Point{
				X: 8,
				Y: 8,
			},
			5,
		},
		Spokes: 0,
	}

	fmt.Printf("%#v\n", w)

	w.X = 42
	fmt.Printf("%#v\n", w)

}
