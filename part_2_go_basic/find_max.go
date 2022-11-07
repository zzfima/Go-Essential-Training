package part2

//FindMax return index and value of maximum at int array
func FindMax(numbers []int) (index int, value int) {
	index = 0
	value = numbers[0]

	for i, v := range numbers {
		if v > value {
			index = i
			value = v
		}
	}

	return index, value
}
