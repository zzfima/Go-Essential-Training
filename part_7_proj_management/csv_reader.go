package part7

import "math"

//TestCase value and square root
type TestCase struct {
	value      float64
	squareRoot float64
}

//Calc calculate square root of value filed
func Calc(tc TestCase) float64 {
	return math.Sqrt(tc.value)
}
