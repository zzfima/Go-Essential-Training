package part2

import "fmt"

//CheckFizzBuzzNumber ...
func CheckFizzBuzzNumber(i int) string {
	switch {
	case divideByFive(i) && divideByThree(i):
		return "fizz buzz"
	case divideByFive(i):
		return "buzz"
	case divideByThree(i):
		return "fizz"
	default:
		return fmt.Sprint(i)
	}
}

func divideByFive(i int) bool {
	return i%5 == 0
}

func divideByThree(i int) bool {
	return i%3 == 0
}
