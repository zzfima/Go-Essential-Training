package part3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestAddNumbers TestAddNumbers
func TestAddNumbers(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	DoubleAt(arr, 2)
	require.Equal(t, arr[2], 6)
}
