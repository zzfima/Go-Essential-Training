package part8

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadRequestsFromFile(t *testing.T) {
	r := ReadRequestsFromFile("bank.json")
	require.Equal(t, 44.2, r.Requests[0].Amount)
	require.Equal(t, "alex", r.Requests[0].User)
	require.Equal(t, "worker", r.Requests[0].Type)
}

func TestReadRequestsFromString(t *testing.T) {

	requestsAsJSONString := `
	{"requests": [
		{"user" : "alex", "type": "worker", "amount": 44.2},
		{"user" : "tom", "type": "manager", "amount": 32.2},
		{"user" : "veronica", "type": "designer", "amount": 45.2}
	]}
	`
	r := ReadRequestsFromString(requestsAsJSONString)
	require.Equal(t, 32.2, r.Requests[1].Amount)
	require.Equal(t, "tom", r.Requests[1].User)
	require.Equal(t, "manager", r.Requests[1].Type)
}

func TestWriteRequestsToFile(t *testing.T) {
	requests := Requests{}
	requests.Requests = append(requests.Requests, Request{User: "Gagarin", Type: "Astronaut", Amount: 9999.99})
	requests.Requests = append(requests.Requests, Request{User: "Mendeleev", Type: "Chenmcalist", Amount: 88888.99})
	WriteRequestsToFile("tmp1.json", requests)

	r := ReadRequestsFromFile("tmp1.json")
	require.Equal(t, 9999.99, r.Requests[0].Amount)
	require.Equal(t, "Gagarin", r.Requests[0].User)
	require.Equal(t, "Astronaut", r.Requests[0].Type)
}

func TestHTTPContentLength(t *testing.T) {
	s, e := GetHTTPContentLength()
	require.Nil(t, e)
	require.Equal(t, 273, s)
}
