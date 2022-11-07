package part4

//Square defines square
type Square struct {
	X      int
	Y      int
	Length int
}

//NewSquare Square factory
func NewSquare(x int, y int, length int) *Square {
	newSquare := Square{x, y, length}
	return &newSquare
}

//Area return square area
func (s Square) Area() int {
	return s.Length * s.Length
}

//Move move X, Y od square
func (s *Square) Move(dx int, dy int) {
	(*s).X += dx
	(*s).Y += dy
}
