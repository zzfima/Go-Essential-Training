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

// TestNextAge TestNextAge
func TestNextAge(t *testing.T) {
	n, e := NextAge(4)
	require.Equal(t, 5, n)
	require.Equal(t, nil, e)

	n, e = NextAge(-3)
	require.Equal(t, 0, n)
	require.Equal(t, "Negative age received", e.Error())
}

func TestGetHTTPText(t *testing.T) {
	url := "https://phet-dev.colorado.edu/html/build-an-atom/0.0.0-3/simple-text-only-test-page.html"

	s, e := GetHTTPText(url)
	require.Equal(t, nil, e)
	require.Equal(t, 992, len(s))
}
