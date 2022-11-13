package part8

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadRequestsFromFile(t *testing.T) {
	r := ReadRequestsFromFile()
	require.Equal(t, 44.2, r.Requests[0].Amount)
	require.Equal(t, "alex", r.Requests[0].User)
	require.Equal(t, "worker", r.Requests[0].Type)
}

func TestReadRequestsFromString(t *testing.T) {
	r := ReadRequestsFromString()
	require.Equal(t, 32.2, r.Requests[1].Amount)
	require.Equal(t, "tom", r.Requests[1].User)
	require.Equal(t, "manager", r.Requests[1].Type)
}
