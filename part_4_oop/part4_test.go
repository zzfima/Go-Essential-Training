package part4

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestBudget(t *testing.T) {
	budget1 := Budget{11, 23.4, time.Now().Add(time.Hour * 24 * 7)}
	require.Equal(t, 11, budget1.CampaignID)
	require.Equal(t, 23.4, budget1.Balance, 0.1)
}

func TestBudgetExpires(t *testing.T) {
	timeNow := time.Now().Add(time.Hour * 24 * 7)
	budget1 := Budget{11, 23.4, timeNow}
	budget1.ExtendBalance(2.2)
	require.InDelta(t, 25.6, budget1.Balance, 0.1)
}

func TestSquares(t *testing.T) {
	s := NewSquare(3, 4, 5.0)
	require.Equal(t, 25.0, s.Area())
	s.Move(2, 2)
	require.Equal(t, 5, s.X)
	require.Equal(t, 6, s.Y)
}

func TestCircles(t *testing.T) {
	c := NewCircle(3, 4, 5.0)
	require.Equal(t, math.Pi*5.0*5.0, c.Area())
	c.Move(2, 2)
	require.Equal(t, 5, c.X)
	require.Equal(t, 6, c.Y)
}

func TestShapes(t *testing.T) {
	s := NewSquare(3, 4, 5.0)
	c := NewCircle(3, 4, 5.0)
	shapes := []Shape{s, c}

	require.Equal(t, 25.0, shapes[0].Area())
	require.Equal(t, math.Pi*5.0*5.0, shapes[1].Area())

	shapes[0].Move(1, 2)
	shapes[1].Move(3, 4)
	require.Equal(t, 4, s.X)
	require.Equal(t, 6, s.Y)
	require.Equal(t, 6, c.X)
	require.Equal(t, 8, c.Y)
}

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
