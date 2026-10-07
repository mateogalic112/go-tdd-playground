// Package mathematics will teach us about math lib
package mathematics

import "time"

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	return Point{150, 60}
}
