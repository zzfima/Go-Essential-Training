package part4

//Shape common interface for a shapes
type Shape interface {
	Area() float64
	Move(dx int, dy int)
}
