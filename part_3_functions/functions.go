package part3

//DoubleArrayAt doubling value at place index
func DoubleArrayAt(arr []int, place int) {
	arr[place] *= 2
}

//DoubleValue doubling value
func DoubleValue(v *int) {
	(*v) *= 2
}
