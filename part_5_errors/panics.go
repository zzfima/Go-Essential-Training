package part5

import "fmt"

//GeneratePanic GeneratePanic
func GeneratePanic() {
	panic("oops")
}

//SafeValue safe way get value from array
func SafeValue(arr []int, index int) (res int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	return arr[index], nil
}
