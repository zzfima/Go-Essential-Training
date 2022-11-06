package numbers

import "fmt"

func main() {
	fmt.Println(AddNumbers(1, 4))
}

//AddNumbers add two numbers
func AddNumbers(i1 int, i2 int) int {
	return i1 + i2
}
