package part4

import "math"

//Circle defines Circle
type Circle struct {
	X      int
	Y      int
	Radius float64
}

//NewCircle Circle factory
func NewCircle(x int, y int, radius float64) *Circle {
	newCircle := Circle{x, y, radius}
	return &newCircle
}

//Area return square area
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

//Move move X, Y od square
func (c *Circle) Move(dx int, dy int) {
	(*c).X += dx
	(*c).Y += dy
}
