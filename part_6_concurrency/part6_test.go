package part6

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetContentType(t *testing.T) {
	r1, e1 := GetContentType("https://github.com/")
	require.Equal(t, "text/html; charset=utf-8", r1)
	require.Nil(t, e1)
}

func TestSerialGetContentType(t *testing.T) {
	startTime := time.Now()
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
	fmt.Println("time passed:", time.Since(startTime))
}

func TestSerialGetContentTypeConcurrent(t *testing.T) {
	startTime := time.Now()
	sites := []string{
		"https://github.com/",
		"https://www.linkedin.com/",
		"https://marketplace.visualstudio.com/",
	}

	r1, e1 := SerialGetContentTypeConcurrent(sites)
	require.Equal(t, "text/html; charset=utf-8", r1[0])
	require.Equal(t, "text/html; charset=utf-8", r1[1])
	require.Equal(t, "text/html; charset=utf-8", r1[2])
	require.Nil(t, e1)
	fmt.Println("time passed:", time.Since(startTime))
}

func TestChannelSingleRun(t *testing.T) {
	i := ChannelSingleRun()
	require.Equal(t, 5, i)
}

func TestChannelMultipleRun(t *testing.T) {
	i := ChannelMultipleRun()
	require.Equal(t, 15, i)
}

func TestChannelMultipleRunWithClose(t *testing.T) {
	i := ChannelMultipleRunWithClose()
	require.Equal(t, 25, i)
}

func TestBufferedChannel(t *testing.T) {
	v := BufferedChannel()
	require.Equal(t, 55, v)
}

func TestWaitForAllGetContentType(t *testing.T) {
	startTime := time.Now()
	sites := []string{
		"https://github.com/",
		"https://www.linkedin.com/",
		"https://marketplace.visualstudio.com/",
	}

	r1, e1 := WaitForAllGetContentType(sites)
	require.Equal(t, "text/html; charset=utf-8", r1[0])
	require.Equal(t, "text/html; charset=utf-8", r1[1])
	require.Equal(t, "text/html; charset=utf-8", r1[2])
	require.Nil(t, e1)
	fmt.Println("time passed:", time.Since(startTime))
}

func TestSelectChannel(t *testing.T) {
	v := SelectChannel(5)
	require.Equal(t, 8, v)

	v = SelectChannel(6)
	require.Equal(t, 6, v)
}

func TestSelectChannelTimeout(t *testing.T) {
	v := SelectChannelTimeout()
	require.Equal(t, 12, v)
}

func TestContextManager(t *testing.T) {
	res := ContextManager()
	require.Equal(t, "https://adsЯus.com/19", res[0].AdURL)
	require.Equal(t, 0.05, res[0].Price)

	require.Equal(t, "https://adsЯus.com/default", res[1].AdURL)
	require.Equal(t, 0.02, res[1].Price)
}
