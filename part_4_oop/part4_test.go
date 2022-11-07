package part4

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestBudget(t *testing.T) {
	budget1 := Budget{11, 23.4, time.Now().Add(time.Hour * 24 * 7)}
	require.Equal(t, 11, budget1.CampaignID)
	require.Equal(t, 23.4, budget1.Balance, 0.1)
}
