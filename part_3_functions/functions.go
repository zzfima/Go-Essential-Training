package part3

import (
	"fmt"
	"io"
	"net/http"
)

// DoubleArrayAt doubling value at place index
func DoubleArrayAt(arr []int, place int) {
	arr[place] *= 2
}

// DoubleValue doubling value
func DoubleValue(v *int) {
	(*v) *= 2
}

// NextAge get next age
func NextAge(currentAge int) (int, error) {
	if currentAge < 0 {
		return 0, fmt.Errorf("Negative age received")
	}
	return currentAge + 1, nil
}

// GetHTTPText get text by url
func GetHTTPText(url string) (string, error) {
	response, e := http.Get(url)
	if e != nil {
		return "", e
	}

	body, e := io.ReadAll(response.Body)
	defer response.Body.Close()
	if e != nil {
		return "", e
	}

	return string(body), nil
}
