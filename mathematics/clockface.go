// Package mathematics will teach us about math lib
package mathematics

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	p := SecondHandPoint(t)
	p = Point{p.X * 90, p.Y * 90}
	p = Point{p.X, -p.Y}
	p = Point{p.X + 150, p.Y + 150}
	return p
}

func SecondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

func SecondHandPoint(t time.Time) Point {
	angle := SecondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
