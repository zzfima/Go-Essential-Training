package part6

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetContentType(t *testing.T) {
	r1, e1 := GetContentType("https://github.com/")
	require.Equal(t, "text/html; charset=utf-8", r1)
	require.Nil(t, e1)
}

func TestSerialGetContentType(t *testing.T) {
	sites := []string{
		"https://github.com/",
		"https://www.linkedin.com/",
		"https://marketplace.visualstudio.com/",
	}

	r1, e1 := SerialGetContentType(sites)
	require.Equal(t, "text/html; charset=utf-8", r1[0])
	require.Equal(t, "text/html; charset=utf-8", r1[1])
	require.Equal(t, "text/html; charset=utf-8", r1[2])
	require.Nil(t, e1)
}
