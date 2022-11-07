package part3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestDoubleArrayAt TestDoubleArrayAt
func TestDoubleArrayAt(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	DoubleArrayAt(arr, 2)
	require.Equal(t, arr[2], 6)
}

// TestDoubleValue TestDoubleValue
func TestDoubleValue(t *testing.T) {
	v := 4
	DoubleValue(&v)
	require.Equal(t, v, 8)
}
