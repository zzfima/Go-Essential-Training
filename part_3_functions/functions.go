package part3

import "fmt"

//DoubleArrayAt doubling value at place index
func DoubleArrayAt(arr []int, place int) {
	arr[place] *= 2
}

//DoubleValue doubling value
func DoubleValue(v *int) {
	(*v) *= 2
}

//NextAge get next age
func NextAge(currentAge int) (nextAge int, e error) {
	if currentAge < 0 {
		return 0, fmt.Errorf("Negative age received")
	}
	return currentAge + 1, nil
}
