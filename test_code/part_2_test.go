package main

import (
	"testing"

	"github.com/stretchr/testify/require"

	numbers "go_essential_training/part_2_go_basic"
)

func TestAddNumbers(t *testing.T) {
	require.Equal(t, numbers.AddNumbers(6, 3), 9)
}
