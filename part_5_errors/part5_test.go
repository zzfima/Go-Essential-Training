package part5

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGeneratePanic(t *testing.T) {
	require.Panics(t, GeneratePanic)
}

func TestSafeValue(t *testing.T) {
	arr := []int{3, 4, 5}

	r1, e1 := SafeValue(arr, 2)
	require.Equal(t, 5, r1)
	require.Nil(t, e1)

	r2, e2 := SafeValue(arr, 3)
	require.Equal(t, 0, r2)
	require.NotNil(t, e2)
}
