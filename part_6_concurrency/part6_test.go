package part6

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetContentTypeType(t *testing.T) {
	r1, e1 := GetContentTypeType("https://github.com/")
	require.Equal(t, "text/html; charset=utf-8", r1)
	require.Nil(t, e1)
}
