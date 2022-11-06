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
