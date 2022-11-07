package part2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddNumbers(t *testing.T) {
	require.Equal(t, 9, 9)
}

func TestFizzBuzz(t *testing.T) {
	require.Equal(t, CheckFizzBuzzNumber(4), "4")
	require.Equal(t, CheckFizzBuzzNumber(15), "fizz buzz")
	require.Equal(t, CheckFizzBuzzNumber(10), "buzz")
	require.Equal(t, CheckFizzBuzzNumber(9), "fizz")
}

func TestFindMax(t *testing.T) {
	arr1 := []int{4, 5, 6, 7, 8, 1, 44, 2, -1}
	arr2 := []int{-222, 44, 55, -3, 10000, -20, 444}

	arr1index, arr1value := FindMax(arr1)
	require.Equal(t, arr1index, 6)
	require.Equal(t, arr1value, 44)

	arr2index, arr2value := FindMax(arr2)
	require.Equal(t, arr2index, 4)
	require.Equal(t, arr2value, 10000)
}

func TestCountWords(t *testing.T) {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`
	expected := make(map[string]int32)
	expected["Needles"] = 2
	expected["Sew"] = 1
	expected["To"] = 1
	expected["a"] = 1
	expected["and"] = 2
	expected["catch"] = 1
	expected["me"] = 2
	expected["pins"] = 2
	expected["sail"] = 1
	expected["the"] = 1
	expected["wind"] = 1

	require.Equal(t, CountWords(text), expected)
}
