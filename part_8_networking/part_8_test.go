package part8

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestXxx(t *testing.T) {
	r := ReadJSON()
	require.Equal(t, 44.2, r.Requests[0].Amount)
	require.Equal(t, "alex", r.Requests[0].User)
	require.Equal(t, "worker", r.Requests[0].Type)
}
